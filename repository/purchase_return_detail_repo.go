package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	PurchaseReturnDetailRepository interface{
		Create(ctx echo.Context, db *gorm.DB, purchaseReturnDetail *domain.PurchaseReturnDetail) (domain.PurchaseReturnDetail, error)
		Update(ctx echo.Context, db *gorm.DB, purchaseReturnDetail *domain.PurchaseReturnDetail) (domain.PurchaseReturnDetail, error)
		Delete(ctx echo.Context, db *gorm.DB, purchaseReturnDetail *domain.PurchaseReturnDetail) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.PurchaseReturnDetail, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.PurchaseReturnDetail, error)
	}

	PurchaseReturnDetailRepositoryImpl struct {

	}
)

func NewPurchaseReturnDetailRepository() PurchaseReturnDetailRepository {
	return &PurchaseReturnDetailRepositoryImpl{}
}

func (repository PurchaseReturnDetailRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, purchaseReturnDetail *domain.PurchaseReturnDetail) (domain.PurchaseReturnDetail, error) {
	db.Create(&purchaseReturnDetail)
	purchaseReturnDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(purchaseReturnDetail.Id))
	return purchaseReturnDetailRes, nil
}

func (repository PurchaseReturnDetailRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, purchaseReturnDetail *domain.PurchaseReturnDetail) (domain.PurchaseReturnDetail, error) {
	db.Where("id = ?", purchaseReturnDetail.Id).Updates(&purchaseReturnDetail)
	purchaseReturnDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(purchaseReturnDetail.Id))
	return purchaseReturnDetailRes, nil
}

func (repository PurchaseReturnDetailRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, purchaseReturnDetail *domain.PurchaseReturnDetail) (bool, error) {
	results := db.Where("id = ?", purchaseReturnDetail.Id).Delete(&purchaseReturnDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|purchaseReturnDetail tidak ditemukan")
	}
	return true, nil
}

func (repository PurchaseReturnDetailRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (purchaseReturnDetailRes domain.PurchaseReturnDetail, err error) {
	results := db.Where(key+" = ?", value).First(&purchaseReturnDetailRes)
	if results.RowsAffected < 1 {
		return purchaseReturnDetailRes, errors.New("NOT_FOUND|purchaseReturnDetail tidak ditemukan")
	}
	return purchaseReturnDetailRes, nil
}

func (repository PurchaseReturnDetailRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (purchaseReturnDetailRes []domain.PurchaseReturnDetail, err error) {
	db.Find(&purchaseReturnDetailRes)
	return purchaseReturnDetailRes, nil
}

