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
	SalesReturnRepository interface{
		Create(ctx echo.Context, db *gorm.DB, salesReturn *domain.SalesReturn) (domain.SalesReturn, error)
		Update(ctx echo.Context, db *gorm.DB, salesReturn *domain.SalesReturn) (domain.SalesReturn, error)
		Delete(ctx echo.Context, db *gorm.DB, salesReturn *domain.SalesReturn) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.SalesReturn, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.SalesReturn, error)
		ReportDatatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.SalesReturnDatatable, int64, int64, error)
	}

	SalesReturnRepositoryImpl struct {

	}
)

func NewSalesReturnRepository() SalesReturnRepository {
	return &SalesReturnRepositoryImpl{}
}

func (repository SalesReturnRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, salesReturn *domain.SalesReturn) (domain.SalesReturn, error) {
	db.Create(&salesReturn)
	salesReturnRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(salesReturn.Id))
	return salesReturnRes, nil
}

func (repository SalesReturnRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, salesReturn *domain.SalesReturn) (domain.SalesReturn, error) {
	db.Where("id = ?", salesReturn.Id).Updates(&salesReturn)
	salesReturnRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(salesReturn.Id))
	return salesReturnRes, nil
}

func (repository SalesReturnRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, salesReturn *domain.SalesReturn) (bool, error) {
	results := db.Where("id = ?", salesReturn.Id).Delete(&salesReturn)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|salesReturn tidak ditemukan")
	}
	return true, nil
}

func (repository SalesReturnRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (salesReturnRes domain.SalesReturn, err error) {
	results := db.Where(key+" = ?", value).First(&salesReturnRes)
	if results.RowsAffected < 1 {
		return salesReturnRes, errors.New("NOT_FOUND|salesReturn tidak ditemukan")
	}
	return salesReturnRes, nil
}

func (repository SalesReturnRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (salesReturnRes []domain.SalesReturn, err error) {
	db.Find(&salesReturnRes)
	return salesReturnRes, nil
}

func (repository SalesReturnRepositoryImpl) ReportDatatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.SalesReturnDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("sales_returns").
		Select(`
			sales_returns.*,
			store_consignments.id store_consignment_id, store_consignments.name store_consignment_name, 
			customers.id customer_id, customers.name customer_name
		`).
		Joins(`
			left join store_consignments on store_consignments.id = sales_returns.store_id
			left join customers on customers.id = sales_returns.customer_id
		`)

	qry.Count(&totalData)
	if search != "" {
		qry.Where("(sales_returns.id = ? OR sales_returns.number LIKE ?)", search, "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("sales_returns.id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}