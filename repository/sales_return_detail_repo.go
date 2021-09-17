package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	SalesReturnDetailRepository interface{
		Create(ctx echo.Context, db *gorm.DB, salesReturnDetail *domain.SalesReturnDetail) (domain.SalesReturnDetail, error)
		Update(ctx echo.Context, db *gorm.DB, salesReturnDetail *domain.SalesReturnDetail) (domain.SalesReturnDetail, error)
		Delete(ctx echo.Context, db *gorm.DB, salesReturnDetail *domain.SalesReturnDetail) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.SalesReturnDetail, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.SalesReturnDetail, error)
	}

	SalesReturnDetailRepositoryImpl struct {

	}
)

func NewSalesReturnDetailRepository() SalesReturnDetailRepository {
	return &SalesReturnDetailRepositoryImpl{}
}

func (repository SalesReturnDetailRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, salesReturnDetail *domain.SalesReturnDetail) (domain.SalesReturnDetail, error) {
	db.Create(&salesReturnDetail)
	salesReturnDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(salesReturnDetail.Id))
	return salesReturnDetailRes, nil
}

func (repository SalesReturnDetailRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, salesReturnDetail *domain.SalesReturnDetail) (domain.SalesReturnDetail, error) {
	db.Where("id = ?", salesReturnDetail.Id).Updates(&salesReturnDetail)
	salesReturnDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(salesReturnDetail.Id))
	return salesReturnDetailRes, nil
}

func (repository SalesReturnDetailRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, salesReturnDetail *domain.SalesReturnDetail) (bool, error) {
	results := db.Where("id = ?", salesReturnDetail.Id).Delete(&salesReturnDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|salesReturnDetail tidak ditemukan")
	}
	return true, nil
}

func (repository SalesReturnDetailRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (salesReturnDetailRes domain.SalesReturnDetail, err error) {
	results := db.Where(key+" = ?", value).First(&salesReturnDetailRes)
	if results.RowsAffected < 1 {
		return salesReturnDetailRes, errors.New("NOT_FOUND|salesReturnDetail tidak ditemukan")
	}
	return salesReturnDetailRes, nil
}

func (repository SalesReturnDetailRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (salesReturnDetailRes []domain.SalesReturnDetail, err error) {
	db.Find(&salesReturnDetailRes)
	return salesReturnDetailRes, nil
}

