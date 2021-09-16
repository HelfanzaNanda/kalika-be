package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	PurchaseOrderDetailRepository interface{
		Create(ctx echo.Context, db *gorm.DB, purchaseOrderDetail *domain.PurchaseOrderDetail) (domain.PurchaseOrderDetail, error)
		Update(ctx echo.Context, db *gorm.DB, purchaseOrderDetail *domain.PurchaseOrderDetail) (domain.PurchaseOrderDetail, error)
		Delete(ctx echo.Context, db *gorm.DB, purchaseOrderDetail *domain.PurchaseOrderDetail) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.PurchaseOrderDetail, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.PurchaseOrderDetail, error)
	}

	PurchaseOrderDetailRepositoryImpl struct {

	}
)

func NewPurchaseOrderDetailRepository() PurchaseOrderDetailRepository {
	return &PurchaseOrderDetailRepositoryImpl{}
}

func (repository PurchaseOrderDetailRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, purchaseOrderDetail *domain.PurchaseOrderDetail) (domain.PurchaseOrderDetail, error) {
	db.Create(&purchaseOrderDetail)
	purchaseOrderDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(purchaseOrderDetail.Id))
	return purchaseOrderDetailRes, nil
}

func (repository PurchaseOrderDetailRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, purchaseOrderDetail *domain.PurchaseOrderDetail) (domain.PurchaseOrderDetail, error) {
	db.Where("id = ?", purchaseOrderDetail.Id).Updates(&purchaseOrderDetail)
	purchaseOrderDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(purchaseOrderDetail.Id))
	return purchaseOrderDetailRes, nil
}

func (repository PurchaseOrderDetailRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, purchaseOrderDetail *domain.PurchaseOrderDetail) (bool, error) {
	results := db.Where("id = ?", purchaseOrderDetail.Id).Delete(&purchaseOrderDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|purchaseOrderDetail tidak ditemukan")
	}
	return true, nil
}

func (repository PurchaseOrderDetailRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (purchaseOrderDetailRes domain.PurchaseOrderDetail, err error) {
	results := db.Where(key+" = ?", value).First(&purchaseOrderDetailRes)
	if results.RowsAffected < 1 {
		return purchaseOrderDetailRes, errors.New("NOT_FOUND|purchaseOrderDetail tidak ditemukan")
	}
	return purchaseOrderDetailRes, nil
}

func (repository PurchaseOrderDetailRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (purchaseOrderDetailRes []domain.PurchaseOrderDetail, err error) {
	db.Find(&purchaseOrderDetailRes)
	return purchaseOrderDetailRes, nil
}

