package repository

import (
	"errors"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type (
	PaymentMethodRepository interface{
		Create(ctx echo.Context, db *gorm.DB, paymentMethod *domain.PaymentMethod) (domain.PaymentMethod, error)
		Update(ctx echo.Context, db *gorm.DB, paymentMethod *domain.PaymentMethod) (domain.PaymentMethod, error)
		Delete(ctx echo.Context, db *gorm.DB, paymentMethod *domain.PaymentMethod) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.PaymentMethod, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.PaymentMethod, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.PaymentMethodDatatable, int64, int64, error)
	}

	PaymentMethodRepositoryImpl struct {

	}
)

func NewPaymentMethodRepository() PaymentMethodRepository {
	return &PaymentMethodRepositoryImpl{}
}

func (repository PaymentMethodRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, paymentMethod *domain.PaymentMethod) (domain.PaymentMethod, error) {
	db.Create(&paymentMethod)
	paymentMethodRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(paymentMethod.Id))
	return paymentMethodRes, nil
}

func (repository PaymentMethodRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, paymentMethod *domain.PaymentMethod) (domain.PaymentMethod, error) {
	db.Where("id = ?", paymentMethod.Id).Updates(&paymentMethod)
	paymentMethodRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(paymentMethod.Id))
	return paymentMethodRes, nil
}

func (repository PaymentMethodRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, paymentMethod *domain.PaymentMethod) (bool, error) {
	results := db.Where("id = ?", paymentMethod.Id).Delete(&paymentMethod)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|paymentMethod tidak ditemukan")
	}
	return true, nil
}

func (repository PaymentMethodRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (paymentMethodRes domain.PaymentMethod, err error) {
	results := db.Where(key+" = ?", value).First(&paymentMethodRes)
	if results.RowsAffected < 1 {
		return paymentMethodRes, errors.New("NOT_FOUND|paymentMethod tidak ditemukan")
	}
	return paymentMethodRes, nil
}

func (repository PaymentMethodRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (paymentMethodRes []domain.PaymentMethod, err error) {
	db.Find(&paymentMethodRes)
	return paymentMethodRes, nil
}

func (repository PaymentMethodRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.PaymentMethodDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("payment_methods")
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(id = ? OR name LIKE ?)", search, "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}

