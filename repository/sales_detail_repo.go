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
		Create(ctx echo.Context, db *gorm.DB, salesDetail *web.SalesPosPost) ([]web.SalesDetailPosGet, error)
		Update(ctx echo.Context, db *gorm.DB, salesDetail *domain.SalesDetail) (domain.SalesDetail, error)
		Delete(ctx echo.Context, db *gorm.DB, salesDetail *domain.SalesDetail) (bool, error)
		DeleteBySales(ctx echo.Context, db *gorm.DB, salesId int) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.SalesDetail, error)
		FindAll(ctx echo.Context, db *gorm.DB, params map[string][]string) ([]web.SalesDetailPosGet, error)
	}

	SalesDetailRepositoryImpl struct {

	}
)

func NewSalesDetailRepository() SalesDetailRepository {
	return &SalesDetailRepositoryImpl{}
}

func (repository SalesDetailRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, salesDetail *web.SalesPosPost) (res []web.SalesDetailPosGet, err error) {
	for _, val := range salesDetail.SalesDetails {
		val.SalesId = salesDetail.Id
		val.Total = float64(val.Qty) * float64(val.UnitPrice)
		db.Table("sales_details").Select("sales_id", "product_id", "qty", "discount_percentage", "discount_value", "total", "unit_price").Create(&val)
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

func (repository SalesDetailRepositoryImpl) DeleteBySales(ctx echo.Context, db *gorm.DB, salesId int) (bool, error) {
	salesDetail := domain.SalesDetail{}
	results := db.Where("sales_id = ?", salesId).Delete(&salesDetail)
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

func (repository SalesDetailRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB, params map[string][]string) (salesDetailRes []web.SalesDetailPosGet, err error) {
	results := db.Table("sales_details").Preload("Product")
	for k, v := range params {
		if v[0] != "" && k != "id" {
			results = results.Where(k+" = ?", v[0])
		}
	}

	results.Find(&salesDetailRes)
	return salesDetailRes, nil
}

