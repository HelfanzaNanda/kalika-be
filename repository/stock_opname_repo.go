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
	StockOpnameRepository interface{
		Create(echo.Context, *gorm.DB, *domain.StockOpname) (web.StockOpnameGet, error)
		Update(echo.Context, *gorm.DB, *domain.StockOpname) (web.StockOpnameGet, error)
		Delete(echo.Context, *gorm.DB, *domain.StockOpname) (bool, error)
		FindById(echo.Context, *gorm.DB, string, string) (web.StockOpnameGet, error)
		FindAll(echo.Context, *gorm.DB) ([]domain.StockOpname, error)
		Datatable(echo.Context, *gorm.DB, string, string, string, string) ([]web.StockOpnameDatatable, int64, int64, error)
	}

	StockOpnameRepositoryImpl struct {

	}
)

func NewStockOpnameRepository() StockOpnameRepository {
	return &StockOpnameRepositoryImpl{}
}

func (repository *StockOpnameRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, stockOpname *domain.StockOpname) (web.StockOpnameGet, error) {
	db.Create(&stockOpname)
	stockOpnameRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(stockOpname.Id))
	return stockOpnameRes, nil
}

func (repository *StockOpnameRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, stockOpname *domain.StockOpname) (web.StockOpnameGet, error) {
	db.Where("id = ?", stockOpname.Id).Updates(&stockOpname)
	stockOpnameRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(stockOpname.Id))
	return stockOpnameRes, nil
}

func (repository *StockOpnameRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, stockOpname *domain.StockOpname) (bool, error) {
	results := db.Where("id = ?", stockOpname.Id).Delete(&stockOpname)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|stockOpname tidak ditemukan")
	}
	return true, nil
}

func (repository *StockOpnameRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (stockOpnameRes web.StockOpnameGet, err error) {
	results := db.Where(key+" = ?", value).First(&stockOpnameRes.StockOpname)
	db.Where("id = ?", stockOpnameRes.StockOpname.StoreId).First(&stockOpnameRes.Store)
	if results.RowsAffected < 1 {
		return stockOpnameRes, errors.New("NOT_FOUND|stockOpname tidak ditemukan")
	}
	return stockOpnameRes, nil
}

func (repository *StockOpnameRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (stockOpnameRes []domain.StockOpname, err error) {
	db.Find(&stockOpnameRes)
	return stockOpnameRes, nil
}

func (repository *StockOpnameRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (recipeRes []web.StockOpnameDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("stock_opnames").Select("stock_opnames.*, users.name as created_by_name")
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(stock_opnames.number LIKE ?)", "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Joins("JOIN users ON users.id = stock_opnames.created_by")
	qry.Order("stock_opnames.id desc")
	qry.Preload("Store").Find(&recipeRes)
	return recipeRes, totalData, totalFiltered, nil
}