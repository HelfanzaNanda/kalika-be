package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	PurchaseOrderDeliveryDetailRepository interface{
		Create(ctx echo.Context, db *gorm.DB, purchaseOrderDeliveryDetail *domain.PurchaseOrderDeliveryDetail) (domain.PurchaseOrderDeliveryDetail, error)
		Update(ctx echo.Context, db *gorm.DB, purchaseOrderDeliveryDetail *domain.PurchaseOrderDeliveryDetail) (domain.PurchaseOrderDeliveryDetail, error)
		Delete(ctx echo.Context, db *gorm.DB, purchaseOrderDeliveryDetail *domain.PurchaseOrderDeliveryDetail) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.PurchaseOrderDeliveryDetail, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.PurchaseOrderDeliveryDetail, error)
	}

	PurchaseOrderDeliveryDetailRepositoryImpl struct {

	}
)

func NewPurchaseOrderDeliveryDetailRepository() PurchaseOrderDeliveryDetailRepository {
	return &PurchaseOrderDeliveryDetailRepositoryImpl{}
}

func (repository PurchaseOrderDeliveryDetailRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, purchaseOrderDeliveryDetail *domain.PurchaseOrderDeliveryDetail) (domain.PurchaseOrderDeliveryDetail, error) {
	db.Create(&purchaseOrderDeliveryDetail)
	purchaseOrderDeliveryDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(purchaseOrderDeliveryDetail.Id))
	return purchaseOrderDeliveryDetailRes, nil
}

func (repository PurchaseOrderDeliveryDetailRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, purchaseOrderDeliveryDetail *domain.PurchaseOrderDeliveryDetail) (domain.PurchaseOrderDeliveryDetail, error) {
	db.Where("id = ?", purchaseOrderDeliveryDetail.Id).Updates(&purchaseOrderDeliveryDetail)
	purchaseOrderDeliveryDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(purchaseOrderDeliveryDetail.Id))
	return purchaseOrderDeliveryDetailRes, nil
}

func (repository PurchaseOrderDeliveryDetailRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, purchaseOrderDeliveryDetail *domain.PurchaseOrderDeliveryDetail) (bool, error) {
	results := db.Where("id = ?", purchaseOrderDeliveryDetail.Id).Delete(&purchaseOrderDeliveryDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|purchaseOrderDeliveryDetail tidak ditemukan")
	}
	return true, nil
}

func (repository PurchaseOrderDeliveryDetailRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (purchaseOrderDeliveryDetailRes domain.PurchaseOrderDeliveryDetail, err error) {
	results := db.Where(key+" = ?", value).First(&purchaseOrderDeliveryDetailRes)
	if results.RowsAffected < 1 {
		return purchaseOrderDeliveryDetailRes, errors.New("NOT_FOUND|purchaseOrderDeliveryDetail tidak ditemukan")
	}
	return purchaseOrderDeliveryDetailRes, nil
}

func (repository PurchaseOrderDeliveryDetailRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (purchaseOrderDeliveryDetailRes []domain.PurchaseOrderDeliveryDetail, err error) {
	db.Find(&purchaseOrderDeliveryDetailRes)
	return purchaseOrderDeliveryDetailRes, nil
}

