package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	PurchaseReturnRepository interface{
		Create(ctx echo.Context, db *gorm.DB, purchaseReturn *domain.PurchaseReturn) (domain.PurchaseReturn, error)
		Update(ctx echo.Context, db *gorm.DB, purchaseReturn *domain.PurchaseReturn) (domain.PurchaseReturn, error)
		Delete(ctx echo.Context, db *gorm.DB, purchaseReturn *domain.PurchaseReturn) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.PurchaseReturn, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.PurchaseReturn, error)
	}

	PurchaseReturnRepositoryImpl struct {

	}
)

func NewPurchaseReturnRepository() PurchaseReturnRepository {
	return &PurchaseReturnRepositoryImpl{}
}

func (repository PurchaseReturnRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, purchaseReturn *domain.PurchaseReturn) (domain.PurchaseReturn, error) {
	db.Create(&purchaseReturn)
	purchaseReturnRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(purchaseReturn.Id))
	return purchaseReturnRes, nil
}

func (repository PurchaseReturnRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, purchaseReturn *domain.PurchaseReturn) (domain.PurchaseReturn, error) {
	db.Where("id = ?", purchaseReturn.Id).Updates(&purchaseReturn)
	purchaseReturnRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(purchaseReturn.Id))
	return purchaseReturnRes, nil
}

func (repository PurchaseReturnRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, purchaseReturn *domain.PurchaseReturn) (bool, error) {
	results := db.Where("id = ?", purchaseReturn.Id).Delete(&purchaseReturn)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|purchaseReturn tidak ditemukan")
	}
	return true, nil
}

func (repository PurchaseReturnRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (purchaseReturnRes domain.PurchaseReturn, err error) {
	results := db.Where(key+" = ?", value).First(&purchaseReturnRes)
	if results.RowsAffected < 1 {
		return purchaseReturnRes, errors.New("NOT_FOUND|purchaseReturn tidak ditemukan")
	}
	return purchaseReturnRes, nil
}

func (repository PurchaseReturnRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (purchaseReturnRes []domain.PurchaseReturn, err error) {
	db.Find(&purchaseReturnRes)
	return purchaseReturnRes, nil
}

