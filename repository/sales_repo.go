package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	SalesRepository interface{
		Create(ctx echo.Context, db *gorm.DB, sales *domain.Sale) (domain.Sale, error)
		Update(ctx echo.Context, db *gorm.DB, sales *domain.Sale) (domain.Sale, error)
		Delete(ctx echo.Context, db *gorm.DB, sales *domain.Sale) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Sale, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Sale, error)
	}

	SalesRepositoryImpl struct {

	}
)

func NewSalesRepository() SalesRepository {
	return &SalesRepositoryImpl{}
}

func (repository SalesRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, sales *domain.Sale) (domain.Sale, error) {
	db.Create(&sales)
	salesRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(sales.Id))
	return salesRes, nil
}

func (repository SalesRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, sales *domain.Sale) (domain.Sale, error) {
	db.Where("id = ?", sales.Id).Updates(&sales)
	salesRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(sales.Id))
	return salesRes, nil
}

func (repository SalesRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, sales *domain.Sale) (bool, error) {
	results := db.Where("id = ?", sales.Id).Delete(&sales)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|sales tidak ditemukan")
	}
	return true, nil
}

func (repository SalesRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (salesRes domain.Sale, err error) {
	results := db.Where(key+" = ?", value).First(&salesRes)
	if results.RowsAffected < 1 {
		return salesRes, errors.New("NOT_FOUND|sales tidak ditemukan")
	}
	return salesRes, nil
}

func (repository SalesRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (salesRes []domain.Sale, err error) {
	db.Find(&salesRes)
	return salesRes, nil
}

