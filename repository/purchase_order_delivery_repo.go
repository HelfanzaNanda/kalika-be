package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	PurchaseOrderDeliveryRepository interface{
		Create(ctx echo.Context, db *gorm.DB, purchaseOrderDelivery *domain.PurchaseOrderDelivery) (domain.PurchaseOrderDelivery, error)
		Update(ctx echo.Context, db *gorm.DB, purchaseOrderDelivery *domain.PurchaseOrderDelivery) (domain.PurchaseOrderDelivery, error)
		Delete(ctx echo.Context, db *gorm.DB, purchaseOrderDelivery *domain.PurchaseOrderDelivery) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.PurchaseOrderDelivery, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.PurchaseOrderDelivery, error)
	}

	PurchaseOrderDeliveryRepositoryImpl struct {

	}
)

func NewPurchaseOrderDeliveryRepository() PurchaseOrderDeliveryRepository {
	return &PurchaseOrderDeliveryRepositoryImpl{}
}

func (repository PurchaseOrderDeliveryRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, purchaseOrderDelivery *domain.PurchaseOrderDelivery) (domain.PurchaseOrderDelivery, error) {
	db.Create(&purchaseOrderDelivery)
	purchaseOrderDeliveryRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(purchaseOrderDelivery.Id))
	return purchaseOrderDeliveryRes, nil
}

func (repository PurchaseOrderDeliveryRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, purchaseOrderDelivery *domain.PurchaseOrderDelivery) (domain.PurchaseOrderDelivery, error) {
	db.Where("id = ?", purchaseOrderDelivery.Id).Updates(&purchaseOrderDelivery)
	purchaseOrderDeliveryRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(purchaseOrderDelivery.Id))
	return purchaseOrderDeliveryRes, nil
}

func (repository PurchaseOrderDeliveryRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, purchaseOrderDelivery *domain.PurchaseOrderDelivery) (bool, error) {
	results := db.Where("id = ?", purchaseOrderDelivery.Id).Delete(&purchaseOrderDelivery)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|purchaseOrderDelivery tidak ditemukan")
	}
	return true, nil
}

func (repository PurchaseOrderDeliveryRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (purchaseOrderDeliveryRes domain.PurchaseOrderDelivery, err error) {
	results := db.Where(key+" = ?", value).First(&purchaseOrderDeliveryRes)
	if results.RowsAffected < 1 {
		return purchaseOrderDeliveryRes, errors.New("NOT_FOUND|purchaseOrderDelivery tidak ditemukan")
	}
	return purchaseOrderDeliveryRes, nil
}

func (repository PurchaseOrderDeliveryRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (purchaseOrderDeliveryRes []domain.PurchaseOrderDelivery, err error) {
	db.Find(&purchaseOrderDeliveryRes)
	return purchaseOrderDeliveryRes, nil
}

