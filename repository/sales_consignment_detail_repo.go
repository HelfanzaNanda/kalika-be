package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	SalesConsignmentDetailRepository interface{
		Create(ctx echo.Context, db *gorm.DB, salesConsignmentDetail *domain.SalesConsignmentDetail) (domain.SalesConsignmentDetail, error)
		Update(ctx echo.Context, db *gorm.DB, salesConsignmentDetail *domain.SalesConsignmentDetail) (domain.SalesConsignmentDetail, error)
		Delete(ctx echo.Context, db *gorm.DB, salesConsignmentDetail *domain.SalesConsignmentDetail) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.SalesConsignmentDetail, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.SalesConsignmentDetail, error)
	}

	SalesConsignmentDetailRepositoryImpl struct {

	}
)

func NewSalesConsignmentDetailRepository() SalesConsignmentDetailRepository {
	return &SalesConsignmentDetailRepositoryImpl{}
}

func (repository SalesConsignmentDetailRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, salesConsignmentDetail *domain.SalesConsignmentDetail) (domain.SalesConsignmentDetail, error) {
	db.Create(&salesConsignmentDetail)
	salesConsignmentDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(salesConsignmentDetail.Id))
	return salesConsignmentDetailRes, nil
}

func (repository SalesConsignmentDetailRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, salesConsignmentDetail *domain.SalesConsignmentDetail) (domain.SalesConsignmentDetail, error) {
	db.Where("id = ?", salesConsignmentDetail.Id).Updates(&salesConsignmentDetail)
	salesConsignmentDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(salesConsignmentDetail.Id))
	return salesConsignmentDetailRes, nil
}

func (repository SalesConsignmentDetailRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, salesConsignmentDetail *domain.SalesConsignmentDetail) (bool, error) {
	results := db.Where("id = ?", salesConsignmentDetail.Id).Delete(&salesConsignmentDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|salesConsignmentDetail tidak ditemukan")
	}
	return true, nil
}

func (repository SalesConsignmentDetailRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (salesConsignmentDetailRes domain.SalesConsignmentDetail, err error) {
	results := db.Where(key+" = ?", value).First(&salesConsignmentDetailRes)
	if results.RowsAffected < 1 {
		return salesConsignmentDetailRes, errors.New("NOT_FOUND|salesConsignmentDetail tidak ditemukan")
	}
	return salesConsignmentDetailRes, nil
}

func (repository SalesConsignmentDetailRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (salesConsignmentDetailRes []domain.SalesConsignmentDetail, err error) {
	db.Find(&salesConsignmentDetailRes)
	return salesConsignmentDetailRes, nil
}

