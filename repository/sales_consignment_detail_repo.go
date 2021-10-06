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
	SalesConsignmentDetailRepository interface{
		Create(ctx echo.Context, db *gorm.DB, salesConsignment *web.SalesConsignmentPost) ([]web.SalesConsignmentDetailGet, error)
		Update(ctx echo.Context, db *gorm.DB, salesConsignmentDetail *domain.SalesConsignmentDetail) (domain.SalesConsignmentDetail, error)
		Delete(ctx echo.Context, db *gorm.DB, salesConsignmentDetail *domain.SalesConsignmentDetail) (bool, error)
		DeleteBySalesConsignment(ctx echo.Context, db *gorm.DB, salesConsignmentId int) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.SalesConsignmentDetail, error)
		FindAll(ctx echo.Context, db *gorm.DB, params map[string][]string) ([]web.SalesConsignmentDetailGet, error)
	}

	SalesConsignmentDetailRepositoryImpl struct {

	}
)

func NewSalesConsignmentDetailRepository() SalesConsignmentDetailRepository {
	return &SalesConsignmentDetailRepositoryImpl{}
}

func (repository *SalesConsignmentDetailRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, salesConsignment *web.SalesConsignmentPost) (res []web.SalesConsignmentDetailGet, err error) {
	for _, val := range salesConsignment.SalesConsignmentDetails {
		val.SalesConsignmentId = salesConsignment.Id
		val.Total = val.Qty * val.UnitPrice
		db.Table("sales_consignment_details").Select("sales_consignment_id", "qty", "product_id", "total", "discount", "unit_price").Create(&val)
		res = append(res, val)
	}

	return res, nil
}

func (repository *SalesConsignmentDetailRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, salesConsignmentDetail *domain.SalesConsignmentDetail) (domain.SalesConsignmentDetail, error) {
	db.Where("id = ?", salesConsignmentDetail.Id).Updates(&salesConsignmentDetail)
	salesConsignmentDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(salesConsignmentDetail.Id))
	return salesConsignmentDetailRes, nil
}

func (repository *SalesConsignmentDetailRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, salesConsignmentDetail *domain.SalesConsignmentDetail) (bool, error) {
	results := db.Where("id = ?", salesConsignmentDetail.Id).Delete(&salesConsignmentDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|salesConsignmentDetail tidak ditemukan")
	}
	return true, nil
}

func (repository *SalesConsignmentDetailRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (salesConsignmentDetailRes domain.SalesConsignmentDetail, err error) {
	results := db.Where(key+" = ?", value).First(&salesConsignmentDetailRes)
	if results.RowsAffected < 1 {
		return salesConsignmentDetailRes, errors.New("NOT_FOUND|salesConsignmentDetail tidak ditemukan")
	}
	return salesConsignmentDetailRes, nil
}

func (repository *SalesConsignmentDetailRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB, params map[string][]string) (salesConsignmentDetailRes []web.SalesConsignmentDetailGet, err error) {
	results := db.Table("sales_consignment_details").Preload("Product")
	for k, v := range params {
		if v[0] != "" && k != "id" {
			results = results.Where(k+" = ?", v[0])
		}
	}

	results.Find(&salesConsignmentDetailRes)
	return salesConsignmentDetailRes, nil
}

func (repository *SalesConsignmentDetailRepositoryImpl) DeleteBySalesConsignment(ctx echo.Context, db *gorm.DB, salesConsignmentId int) (bool, error) {
	salesConsignmentDetail := domain.SalesConsignmentDetail{}
	results := db.Where("sales_consignment_id = ?", salesConsignmentId).Delete(&salesConsignmentDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|salesDetail tidak ditemukan")
	}
	return true, nil
}
