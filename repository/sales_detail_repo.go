package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"
)

type (
	SalesDetailRepository interface{
		Create(ctx echo.Context, db *gorm.DB, salesDetail *web.SalesPosPost) ([]domain.SalesDetail, error)
		Update(ctx echo.Context, db *gorm.DB, salesDetail *domain.SalesDetail) (domain.SalesDetail, error)
		Delete(ctx echo.Context, db *gorm.DB, salesDetail *domain.SalesDetail) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.SalesDetail, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.SalesDetail, error)
	}

	SalesDetailRepositoryImpl struct {

	}
)

func NewSalesDetailRepository() SalesDetailRepository {
	return &SalesDetailRepositoryImpl{}
}

func (repository SalesDetailRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, salesDetail *web.SalesPosPost) (res []domain.SalesDetail, err error) {
	for _, val := range salesDetail.SalesDetails {
		val.SalesId = salesDetail.Id
		val.Total = float64(val.Qty) * float64(val.UnitPrice)
		db.Create(&val)
		res = append(res, val)
	}

	return res, nil
}

func (repository SalesDetailRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, salesDetail *domain.SalesDetail) (domain.SalesDetail, error) {
	db.Where("id = ?", salesDetail.Id).Updates(&salesDetail)
	salesDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(salesDetail.Id))
	return salesDetailRes, nil
}

func (repository SalesDetailRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, salesDetail *domain.SalesDetail) (bool, error) {
	results := db.Where("id = ?", salesDetail.Id).Delete(&salesDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|salesDetail tidak ditemukan")
	}
	return true, nil
}

func (repository SalesDetailRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (salesDetailRes domain.SalesDetail, err error) {
	results := db.Where(key+" = ?", value).First(&salesDetailRes)
	if results.RowsAffected < 1 {
		return salesDetailRes, errors.New("NOT_FOUND|salesDetail tidak ditemukan")
	}
	return salesDetailRes, nil
}

func (repository SalesDetailRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (salesDetailRes []domain.SalesDetail, err error) {
	db.Find(&salesDetailRes)
	return salesDetailRes, nil
}

