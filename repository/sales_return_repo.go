package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	SalesReturnRepository interface{
		Create(ctx echo.Context, db *gorm.DB, salesReturn *domain.SalesReturn) (domain.SalesReturn, error)
		Update(ctx echo.Context, db *gorm.DB, salesReturn *domain.SalesReturn) (domain.SalesReturn, error)
		Delete(ctx echo.Context, db *gorm.DB, salesReturn *domain.SalesReturn) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.SalesReturn, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.SalesReturn, error)
	}

	SalesReturnRepositoryImpl struct {

	}
)

func NewSalesReturnRepository() SalesReturnRepository {
	return &SalesReturnRepositoryImpl{}
}

func (repository SalesReturnRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, salesReturn *domain.SalesReturn) (domain.SalesReturn, error) {
	db.Create(&salesReturn)
	salesReturnRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(salesReturn.Id))
	return salesReturnRes, nil
}

func (repository SalesReturnRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, salesReturn *domain.SalesReturn) (domain.SalesReturn, error) {
	db.Where("id = ?", salesReturn.Id).Updates(&salesReturn)
	salesReturnRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(salesReturn.Id))
	return salesReturnRes, nil
}

func (repository SalesReturnRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, salesReturn *domain.SalesReturn) (bool, error) {
	results := db.Where("id = ?", salesReturn.Id).Delete(&salesReturn)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|salesReturn tidak ditemukan")
	}
	return true, nil
}

func (repository SalesReturnRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (salesReturnRes domain.SalesReturn, err error) {
	results := db.Where(key+" = ?", value).First(&salesReturnRes)
	if results.RowsAffected < 1 {
		return salesReturnRes, errors.New("NOT_FOUND|salesReturn tidak ditemukan")
	}
	return salesReturnRes, nil
}

func (repository SalesReturnRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (salesReturnRes []domain.SalesReturn, err error) {
	db.Find(&salesReturnRes)
	return salesReturnRes, nil
}

