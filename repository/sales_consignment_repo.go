package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	SalesConsignmentRepository interface{
		Create(ctx echo.Context, db *gorm.DB, salesConsignment *domain.SalesConsignment) (domain.SalesConsignment, error)
		Update(ctx echo.Context, db *gorm.DB, salesConsignment *domain.SalesConsignment) (domain.SalesConsignment, error)
		Delete(ctx echo.Context, db *gorm.DB, salesConsignment *domain.SalesConsignment) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.SalesConsignment, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.SalesConsignment, error)
	}

	SalesConsignmentRepositoryImpl struct {

	}
)

func NewSalesConsignmentRepository() SalesConsignmentRepository {
	return &SalesConsignmentRepositoryImpl{}
}

func (repository SalesConsignmentRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, salesConsignment *domain.SalesConsignment) (domain.SalesConsignment, error) {
	db.Create(&salesConsignment)
	salesConsignmentRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(salesConsignment.Id))
	return salesConsignmentRes, nil
}

func (repository SalesConsignmentRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, salesConsignment *domain.SalesConsignment) (domain.SalesConsignment, error) {
	db.Where("id = ?", salesConsignment.Id).Updates(&salesConsignment)
	salesConsignmentRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(salesConsignment.Id))
	return salesConsignmentRes, nil
}

func (repository SalesConsignmentRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, salesConsignment *domain.SalesConsignment) (bool, error) {
	results := db.Where("id = ?", salesConsignment.Id).Delete(&salesConsignment)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|salesConsignment tidak ditemukan")
	}
	return true, nil
}

func (repository SalesConsignmentRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (salesConsignmentRes domain.SalesConsignment, err error) {
	results := db.Where(key+" = ?", value).First(&salesConsignmentRes)
	if results.RowsAffected < 1 {
		return salesConsignmentRes, errors.New("NOT_FOUND|salesConsignment tidak ditemukan")
	}
	return salesConsignmentRes, nil
}

func (repository SalesConsignmentRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (salesConsignmentRes []domain.SalesConsignment, err error) {
	db.Find(&salesConsignmentRes)
	return salesConsignmentRes, nil
}

