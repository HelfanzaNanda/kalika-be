package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"time"
)

type (
	PaymentRepository interface{
		Create(ctx echo.Context, db *gorm.DB, payment *domain.Payment) (domain.Payment, error)
		Update(ctx echo.Context, db *gorm.DB, payment *domain.Payment) (domain.Payment, error)
		Delete(ctx echo.Context, db *gorm.DB, payment *domain.Payment) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string, params map[string][]string) (domain.Payment, error)
		FindAll(ctx echo.Context, db *gorm.DB, params map[string][]string) ([]domain.Payment, error)
	}

	PaymentRepositoryImpl struct {

	}
)

func NewPaymentRepository() PaymentRepository {
	return &PaymentRepositoryImpl{}
}

func (repository PaymentRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, payment *domain.Payment) (domain.Payment, error) {
	payment.Number = "PY"+helpers.IntToString(int(time.Now().Unix()))
	db.Create(&payment)
	paymentRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(payment.Id), map[string][]string{})
	return paymentRes, nil
}

func (repository PaymentRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, payment *domain.Payment) (res domain.Payment, err error) {
	if err := db.Where("id = ?", payment.Id).First(&res).Error; err != nil {
		return res, errors.New("NOT_FOUND|pembayaran tidak ditemukan")
	}

	db.Model(&res).Updates(&payment)

	paymentRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(payment.Id), map[string][]string{})
	return paymentRes, nil
}

func (repository PaymentRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, payment *domain.Payment) (bool, error) {
	results := db.Where("id = ?", payment.Id).Delete(&payment)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|payment tidak ditemukan")
	}
	return true, nil
}

func (repository PaymentRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string, params map[string][]string) (paymentRes domain.Payment, err error) {
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

func (repository PaymentRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB, params map[string][]string) (paymentRes []domain.Payment, err error) {
	db.Find(&paymentRes)
	return paymentRes, nil
}

