package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	PurchaseInvoiceRepository interface{
		Create(ctx echo.Context, db *gorm.DB, purchaseInvoice *domain.PurchaseInvoice) (domain.PurchaseInvoice, error)
		Update(ctx echo.Context, db *gorm.DB, purchaseInvoice *domain.PurchaseInvoice) (domain.PurchaseInvoice, error)
		Delete(ctx echo.Context, db *gorm.DB, purchaseInvoice *domain.PurchaseInvoice) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.PurchaseInvoice, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.PurchaseInvoice, error)
	}

	PurchaseInvoiceRepositoryImpl struct {

	}
)

func NewPurchaseInvoiceRepository() PurchaseInvoiceRepository {
	return &PurchaseInvoiceRepositoryImpl{}
}

func (repository PurchaseInvoiceRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, purchaseInvoice *domain.PurchaseInvoice) (domain.PurchaseInvoice, error) {
	db.Create(&purchaseInvoice)
	purchaseInvoiceRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(purchaseInvoice.Id))
	return purchaseInvoiceRes, nil
}

func (repository PurchaseInvoiceRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, purchaseInvoice *domain.PurchaseInvoice) (domain.PurchaseInvoice, error) {
	db.Where("id = ?", purchaseInvoice.Id).Updates(&purchaseInvoice)
	purchaseInvoiceRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(purchaseInvoice.Id))
	return purchaseInvoiceRes, nil
}

func (repository PurchaseInvoiceRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, purchaseInvoice *domain.PurchaseInvoice) (bool, error) {
	results := db.Where("id = ?", purchaseInvoice.Id).Delete(&purchaseInvoice)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|purchaseInvoice tidak ditemukan")
	}
	return true, nil
}

func (repository PurchaseInvoiceRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (purchaseInvoiceRes domain.PurchaseInvoice, err error) {
	results := db.Where(key+" = ?", value).First(&purchaseInvoiceRes)
	if results.RowsAffected < 1 {
		return purchaseInvoiceRes, errors.New("NOT_FOUND|purchaseInvoice tidak ditemukan")
	}
	return purchaseInvoiceRes, nil
}

func (repository PurchaseInvoiceRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (purchaseInvoiceRes []domain.PurchaseInvoice, err error) {
	db.Find(&purchaseInvoiceRes)
	return purchaseInvoiceRes, nil
}

