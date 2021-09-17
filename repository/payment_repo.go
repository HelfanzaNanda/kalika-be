package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	PaymentRepository interface{
		Create(ctx echo.Context, db *gorm.DB, payment *domain.Payment) (domain.Payment, error)
		Update(ctx echo.Context, db *gorm.DB, payment *domain.Payment) (domain.Payment, error)
		Delete(ctx echo.Context, db *gorm.DB, payment *domain.Payment) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Payment, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Payment, error)
	}

	PaymentRepositoryImpl struct {

	}
)

func NewPaymentRepository() PaymentRepository {
	return &PaymentRepositoryImpl{}
}

func (repository PaymentRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, payment *domain.Payment) (domain.Payment, error) {
	db.Create(&payment)
	paymentRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(payment.Id))
	return paymentRes, nil
}

func (repository PaymentRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, payment *domain.Payment) (domain.Payment, error) {
	db.Where("id = ?", payment.Id).Updates(&payment)
	paymentRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(payment.Id))
	return paymentRes, nil
}

func (repository PaymentRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, payment *domain.Payment) (bool, error) {
	results := db.Where("id = ?", payment.Id).Delete(&payment)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|payment tidak ditemukan")
	}
	return true, nil
}

func (repository PaymentRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (paymentRes domain.Payment, err error) {
	results := db.Where(key+" = ?", value).First(&paymentRes)
	if results.RowsAffected < 1 {
		return paymentRes, errors.New("NOT_FOUND|payment tidak ditemukan")
	}
	return paymentRes, nil
}

func (repository PaymentRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (paymentRes []domain.Payment, err error) {
	db.Find(&paymentRes)
	return paymentRes, nil
}

