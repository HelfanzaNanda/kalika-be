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
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string, filter map[string]string) ([]web.PaymentDatatable, int64, int64, error)
		FindByModel(ctx echo.Context, db *gorm.DB, model string, model_id int, filter map[string]string) (web.PaymentGet, error)
		FindByCreatedAt(ctx echo.Context, db *gorm.DB, filter *web.PaymentReportFilterDatatable) ([]web.PaymentReportGet, error)
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


func (repository *PaymentRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string, filter map[string]string) (datatableRes []web.PaymentDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("payments")
	qry.Select(`
		payments.*,
		payment_methods.name payment_method,
		users.name created_by_name,
		stores.name store_name
	`)
	qry.Joins(`
		left join payment_methods on payment_methods.id = payments.payment_method_id
		left join users on users.id = payments.created_by
		left join stores on stores.id = payments.store_id
	`)
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(payments.number LIKE ? OR stores.name LIKE ? OR payment_methods.name LIKE ? OR users.name LIKE ?)", "%"+search+"%", "%"+search+"%", "%"+search+"%" ,"%"+search+"%")
	}
	if filter["start_date"] != "" && filter["end_date"] != ""{
		qry.Where("(DATE(payments.created_at) BETWEEN ? AND ?)", filter["start_date"], filter["end_date"])
	}
	if filter["store_id"] != "" {
		qry.Where("(payments.store_id = ?)", filter["store_id"])
	}
	if filter["created_by"] != "" {
		qry.Where("(payments.created_by = ?)", filter["created_by"])
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("payments.id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}

func (repository *PaymentRepositoryImpl) FindByModel(ctx echo.Context, db *gorm.DB, model string, model_id int, filter map[string]string) (res web.PaymentGet, err error) {
	qry := db.Table("payments")
	qry.Select("payments.*, payment_methods.name payment_method_name")
	qry.Joins("join payment_methods on payment_methods.id = payments.payment_method_id")
	if model != "" {
		qry.Where("(payments.model = ?)", model)
	}
	if model_id != 0 {
		qry.Where("(payments.model_id = ?)", model_id)
	}
	if filter["payment_method_id"] != "" {
		qry.Where("(payments.payment_method_id = ?)", filter["payment_method_id"])
	}
	qry.First(&res)
	// if qry.RowsAffected < 1 {
	// 	return res, errors.New("NOT_FOUND|payment tidak ditemukan")
	// }
	return res, nil
}

func (repository *PaymentRepositoryImpl) FindByCreatedAt(ctx echo.Context, db *gorm.DB, filter *web.PaymentReportFilterDatatable) (res []web.PaymentReportGet, err error) {
	qry := db.Table("payments")
	qry.Select(`
		payments.*,
		payment_methods.name payment_method,
		users.name created_by_name,
		stores.name store_name
	`)
	qry.Joins(`
		left join payment_methods on payment_methods.id = payments.payment_method_id
		left join users on users.id = payments.created_by
		left join stores on stores.id = payments.store_id
	`)
	if filter.StartDate != "" && filter.EndDate != ""{
		qry.Where("(DATE(payments.created_at) BETWEEN ? AND ?)", filter.StartDate, filter.EndDate)
	}
	if filter.StoreId != 0 {
		qry.Where("(payments.store_id = ?)", filter.StoreId)
	}
	if filter.CreatedBy != 0 {
		qry.Where("(payments.created_by = ?)", filter.CreatedBy)
	}
	qry.Order("payments.id desc")
	qry.Find(&res)
	return res, nil
}

