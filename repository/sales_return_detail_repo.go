package repository

import (
	"errors"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type (
	SalesReturnDetailRepository interface {
		Create(ctx echo.Context, db *gorm.DB, salesReturn *web.SalesReturnPost) (web.SalesReturnPost, error)
		Update(ctx echo.Context, db *gorm.DB, salesReturnDetail *domain.SalesReturnDetail) (domain.SalesReturnDetail, error)
		Delete(ctx echo.Context, db *gorm.DB, salesReturnDetail *domain.SalesReturnDetail) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.SalesReturnDetail, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.SalesReturnDetail, error)
		DeleteBySalesReturnId(ctx echo.Context, db *gorm.DB, salesReturnId int) (bool, error)
		FindBySalesReturnId(ctx echo.Context, db *gorm.DB, salesReturnId int) ([]domain.SalesReturnDetail, error)
	}

	SalesReturnDetailRepositoryImpl struct {
	}
)

func NewSalesReturnDetailRepository() SalesReturnDetailRepository {
	return &SalesReturnDetailRepositoryImpl{}
}

func (repository SalesReturnDetailRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, salesReturn *web.SalesReturnPost) (res web.SalesReturnPost, err error) {
	var total float64 = 0
	for _, val := range salesReturn.SalesReturnDetails {
		if val.Qty > 0 {
			val.SalesReturnId = salesReturn.Id
			db.Create(&val)
			total += val.Total
			res.SalesReturnDetails = append(res.SalesReturnDetails, val)
		}
	}
	res.Total = total
	return res, nil
}

func (repository SalesReturnDetailRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, salesReturnDetail *domain.SalesReturnDetail) (domain.SalesReturnDetail, error) {
	db.Where("id = ?", salesReturnDetail.Id).Updates(&salesReturnDetail)
	salesReturnDetailRes, _ := repository.FindById(ctx, db, "id", helpers.IntToString(salesReturnDetail.Id))
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

func (repository SalesReturnDetailRepositoryImpl) DeleteBySalesReturnId(ctx echo.Context, db *gorm.DB, salesReturnId int) (bool, error) {
	results := db.Where("sales_return_id = ?", salesReturnId).Delete(domain.SalesReturnDetail{})
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|sales return Detail tidak ditemukan")
	}
	return true, nil
}

func (repository SalesReturnDetailRepositoryImpl) FindBySalesReturnId(ctx echo.Context, db *gorm.DB, salesReturnId int) (salesReturnDetailRes []domain.SalesReturnDetail, err error) {
	results := db.Where("sales_return_id = ?", salesReturnId).Find(&salesReturnDetailRes)
	if results.RowsAffected < 1 {
		return salesReturnDetailRes, errors.New("NOT_FOUND|sales return Detail tidak ditemukan")
	}
	return salesReturnDetailRes, nil
}
