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
	StockOpnameDetailRepository interface{
		Create(echo.Context, *gorm.DB, []domain.StockOpnameDetail) ([]web.StockOpnameDetailGet, error)
		Update(echo.Context, *gorm.DB, *domain.StockOpnameDetail) (domain.StockOpnameDetail, error)
		Delete(echo.Context, *gorm.DB, *domain.StockOpnameDetail) (bool, error)
		DeleteByStockOpname(echo.Context, *gorm.DB, int) (bool, error)
		FindById(echo.Context, *gorm.DB, string, string) (domain.StockOpnameDetail, error)
		FindAll(echo.Context, *gorm.DB, map[string][]string) ([]web.StockOpnameDetailGet, error)
	}

	StockOpnameDetailRepositoryImpl struct {

	}
)

func NewStockOpnameDetailRepository() StockOpnameDetailRepository {
	return &StockOpnameDetailRepositoryImpl{}
}

func (repository *StockOpnameDetailRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, stockOpnameDetail []domain.StockOpnameDetail) (res []web.StockOpnameDetailGet, err error) {
	for _, val := range stockOpnameDetail {
		db.Create(&val)
	}

	detailSearch := make(map[string][]string)
	detailSearch["stock_opname_id"] = append(detailSearch["stock_opname_id"], helpers.IntToString(stockOpnameDetail[0].StockOpnameId))

	res, _ = repository.FindAll(ctx, db, detailSearch)

	return res, nil
}

func (repository *StockOpnameDetailRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, stockOpnameDetail *domain.StockOpnameDetail) (domain.StockOpnameDetail, error) {
	db.Where("id = ?", stockOpnameDetail.Id).Updates(&stockOpnameDetail)
	stockOpnameDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(stockOpnameDetail.Id))
	return stockOpnameDetailRes, nil
}

func (repository *StockOpnameDetailRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, stockOpnameDetail *domain.StockOpnameDetail) (bool, error) {
	results := db.Where("id = ?", stockOpnameDetail.Id).Delete(&stockOpnameDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|stockOpnameDetail tidak ditemukan")
	}
	return true, nil
}

func (repository *StockOpnameDetailRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (stockOpnameDetailRes domain.StockOpnameDetail, err error) {
	results := db.Where(key+" = ?", value).First(&stockOpnameDetailRes)
	if results.RowsAffected < 1 {
		return stockOpnameDetailRes, errors.New("NOT_FOUND|stockOpnameDetail tidak ditemukan")
	}
	return stockOpnameDetailRes, nil
}

func (repository *StockOpnameDetailRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB, params map[string][]string) (stockOpnameDetailRes []web.StockOpnameDetailGet, err error) {
	results := db.Table("stock_opname_details").Preload("Product").Preload("Store")
	for k, v := range params {
		if v[0] != "" && k != "id" {
			results = results.Where(k+" = ?", v[0])
		}
	}
	results.Find(&stockOpnameDetailRes)
	return stockOpnameDetailRes, nil
}

func (repository *StockOpnameDetailRepositoryImpl) DeleteByStockOpname(ctx echo.Context, db *gorm.DB, stockOpnameId int) (bool, error) {
	stockOpnameDetail := domain.StockOpnameDetail{}
	results := db.Where("stock_opname_id = ?", stockOpnameId).Delete(&stockOpnameDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|stockOpnameDetail tidak ditemukan")
	}
	return true, nil
}
