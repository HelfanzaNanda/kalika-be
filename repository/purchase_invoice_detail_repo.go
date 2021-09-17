package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	PurchaseInvoiceDetailRepository interface{
		Create(ctx echo.Context, db *gorm.DB, purchaseInvoiceDetail *domain.PurchaseInvoiceDetail) (domain.PurchaseInvoiceDetail, error)
		Update(ctx echo.Context, db *gorm.DB, purchaseInvoiceDetail *domain.PurchaseInvoiceDetail) (domain.PurchaseInvoiceDetail, error)
		Delete(ctx echo.Context, db *gorm.DB, purchaseInvoiceDetail *domain.PurchaseInvoiceDetail) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.PurchaseInvoiceDetail, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.PurchaseInvoiceDetail, error)
	}

	PurchaseInvoiceDetailRepositoryImpl struct {

	}
)

func NewPurchaseInvoiceDetailRepository() PurchaseInvoiceDetailRepository {
	return &PurchaseInvoiceDetailRepositoryImpl{}
}

func (repository PurchaseInvoiceDetailRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, purchaseInvoiceDetail *domain.PurchaseInvoiceDetail) (domain.PurchaseInvoiceDetail, error) {
	db.Create(&purchaseInvoiceDetail)
	purchaseInvoiceDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(purchaseInvoiceDetail.Id))
	return purchaseInvoiceDetailRes, nil
}

func (repository PurchaseInvoiceDetailRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, purchaseInvoiceDetail *domain.PurchaseInvoiceDetail) (domain.PurchaseInvoiceDetail, error) {
	db.Where("id = ?", purchaseInvoiceDetail.Id).Updates(&purchaseInvoiceDetail)
	purchaseInvoiceDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(purchaseInvoiceDetail.Id))
	return purchaseInvoiceDetailRes, nil
}

func (repository PurchaseInvoiceDetailRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, purchaseInvoiceDetail *domain.PurchaseInvoiceDetail) (bool, error) {
	results := db.Where("id = ?", purchaseInvoiceDetail.Id).Delete(&purchaseInvoiceDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|purchaseInvoiceDetail tidak ditemukan")
	}
	return true, nil
}

func (repository PurchaseInvoiceDetailRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (purchaseInvoiceDetailRes domain.PurchaseInvoiceDetail, err error) {
	results := db.Where(key+" = ?", value).First(&purchaseInvoiceDetailRes)
	if results.RowsAffected < 1 {
		return purchaseInvoiceDetailRes, errors.New("NOT_FOUND|purchaseInvoiceDetail tidak ditemukan")
	}
	return purchaseInvoiceDetailRes, nil
}

func (repository PurchaseInvoiceDetailRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (purchaseInvoiceDetailRes []domain.PurchaseInvoiceDetail, err error) {
	db.Find(&purchaseInvoiceDetailRes)
	return purchaseInvoiceDetailRes, nil
}

