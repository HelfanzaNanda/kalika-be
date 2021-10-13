package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"
	"time"
)

type (
	PaymentRepository interface{
		Create(ctx echo.Context, db *gorm.DB, payment *domain.Payment) (domain.Payment, error)
		Update(ctx echo.Context, db *gorm.DB, payment *domain.Payment) (domain.Payment, error)
		Delete(ctx echo.Context, db *gorm.DB, payment *domain.Payment) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string, params map[string][]string) (domain.Payment, error)
		FindAll(ctx echo.Context, db *gorm.DB, params map[string][]string) ([]domain.Payment, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.PaymentDatatable, int64, int64, error)
	}

	PaymentRepositoryImpl struct {

	}
)

func NewPaymentRepository() PaymentRepository {
	return &PaymentRepositoryImpl{}
}

func (repository *PaymentRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, payment *domain.Payment) (domain.Payment, error) {
	payment.Number = "PY"+helpers.IntToString(int(time.Now().Unix()))
	db.Create(&payment)
	paymentRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(payment.Id), map[string][]string{})
	return paymentRes, nil
}

func (repository *PaymentRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, payment *domain.Payment) (res domain.Payment, err error) {
	if err := db.Where("id = ?", payment.Id).First(&res).Error; err != nil {
		return res, errors.New("NOT_FOUND|pembayaran tidak ditemukan")
	}

	db.Model(&res).Updates(&payment)

	paymentRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(payment.Id), map[string][]string{})
	return paymentRes, nil
}

func (repository *PaymentRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, payment *domain.Payment) (bool, error) {
	results := db.Where("id = ?", payment.Id).Delete(&payment)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|payment tidak ditemukan")
	}
	return true, nil
}

func (repository *PaymentRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string, params map[string][]string) (paymentRes domain.Payment, err error) {
	results := db.Where(key+" = ?", value)
	for k, v := range params {
		if v[0] != "" && k != "id" {
			results = results.Where(k+" = ?", v[0])
		}
	}
	results.First(&paymentRes)
	if results.RowsAffected < 1 {
		return paymentRes, errors.New("NOT_FOUND|payment tidak ditemukan")
	}
	return paymentRes, nil
}

func (repository *PaymentRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB, params map[string][]string) (paymentRes []domain.Payment, err error) {
	db.Find(&paymentRes)
	return paymentRes, nil
}

func (repository *PaymentRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.PaymentDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("payments").
		Select(`
		payments.*,
		payment_methods.name payment_method,
		users.name created_by_name,
		stores.name store_name
	`).
		Joins(`
		left join payment_methods on payment_methods.id = payments.payment_method_id
		left join users on users.id = payments.created_by
		left join stores on stores.id = payments.store_id
	`)
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(payments.number LIKE ? OR stores.name LIKE ? OR payment_methods.name LIKE ? OR users.name LIKE ?)", "%"+search+"%", "%"+search+"%", "%"+search+"%" ,"%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("payments.id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}

