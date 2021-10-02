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
	SalesRepository interface{
		Create(ctx echo.Context, db *gorm.DB, sales *domain.Sale) (domain.Sale, error)
		Update(ctx echo.Context, db *gorm.DB, sales *domain.Sale) (domain.Sale, error)
		Delete(ctx echo.Context, db *gorm.DB, sales *domain.Sale) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Sale, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Sale, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.SaleDatatable, int64, int64, error)
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

func (repository SalesRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.SaleDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("sales").
		Select(`
			sales.*,
			stores.id store_id, stores.name store_name, 
			customers.id customer_id, customers.name customer_name, 
			cash_registers.id cash_register_id, cash_registers.cash_in_hand cash_register_cash_in_hand 
		`).
		Joins(`
			left join stores on stores.id = sales.store_id
			left join customers on customers.id = sales.customer_id
			left join cash_registers on cash_registers.id = sales.cash_register_id
		`)

	qry.Count(&totalData)
	if search != "" {
		qry.Where("(sales.id = ? OR sales.number LIKE ?)", search, "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("sales.id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}
