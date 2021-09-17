package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	PurchaseOrderRepository interface{
		Create(ctx echo.Context, db *gorm.DB, purchaseOrder *domain.PurchaseOrder) (domain.PurchaseOrder, error)
		Update(ctx echo.Context, db *gorm.DB, purchaseOrder *domain.PurchaseOrder) (domain.PurchaseOrder, error)
		Delete(ctx echo.Context, db *gorm.DB, purchaseOrder *domain.PurchaseOrder) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.PurchaseOrder, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.PurchaseOrder, error)
	}

	PurchaseOrderRepositoryImpl struct {

	}
)

func NewPurchaseOrderRepository() PurchaseOrderRepository {
	return &PurchaseOrderRepositoryImpl{}
}

func (repository PurchaseOrderRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, purchaseOrder *domain.PurchaseOrder) (domain.PurchaseOrder, error) {
	db.Create(&purchaseOrder)
	purchaseOrderRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(purchaseOrder.Id))
	return purchaseOrderRes, nil
}

func (repository PurchaseOrderRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, purchaseOrder *domain.PurchaseOrder) (domain.PurchaseOrder, error) {
	db.Where("id = ?", purchaseOrder.Id).Updates(&purchaseOrder)
	purchaseOrderRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(purchaseOrder.Id))
	return purchaseOrderRes, nil
}

func (repository PurchaseOrderRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, purchaseOrder *domain.PurchaseOrder) (bool, error) {
	results := db.Where("id = ?", purchaseOrder.Id).Delete(&purchaseOrder)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|purchaseOrder tidak ditemukan")
	}
	return true, nil
}

func (repository PurchaseOrderRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (purchaseOrderRes domain.PurchaseOrder, err error) {
	results := db.Where(key+" = ?", value).First(&purchaseOrderRes)
	if results.RowsAffected < 1 {
		return purchaseOrderRes, errors.New("NOT_FOUND|purchaseOrder tidak ditemukan")
	}
	return purchaseOrderRes, nil
}

func (repository PurchaseOrderRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (purchaseOrderRes []domain.PurchaseOrder, err error) {
	db.Find(&purchaseOrderRes)
	return purchaseOrderRes, nil
}

