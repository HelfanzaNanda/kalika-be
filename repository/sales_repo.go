package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"
	"time"
)

type (
	SalesRepository interface{
		Create(ctx echo.Context, db *gorm.DB, sales *web.SalesPosPost) (domain.Sale, error)
		Update(ctx echo.Context, db *gorm.DB, sales *domain.Sale) (domain.Sale, error)
		Delete(ctx echo.Context, db *gorm.DB, sales *domain.Sale) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Sale, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Sale, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.SaleDatatable, int64, int64, error)
		ReportDatatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string, filter map[string]string) ([]web.SaleDatatable, int64, int64, error)
		FindByCreatedAt(ctx echo.Context, db *gorm.DB, dateRange *web.DateRange) ([]web.SalesPosGet, error)
	}

	SalesRepositoryImpl struct {

	}
)

func NewSalesRepository() SalesRepository {
	return &SalesRepositoryImpl{}
}

func (repository SalesRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, sales *web.SalesPosPost) (domain.Sale, error) {
	m := domain.Sale{}

	m.Number = "INV"+helpers.IntToString(int(time.Now().Unix()))
	m.StoreId = helpers.StringToInt(ctx.Get("userInfo").(map[string]interface{})["store_id"].(string))
	m.CreatedBy = helpers.StringToInt(ctx.Get("userInfo").(map[string]interface{})["id"].(string))
	m.CustomerPay = sales.CustomerPay
	m.CustomerChange = sales.CustomerChange
	m.Total = sales.Total
	m.DiscountValue = sales.DiscountValue
	m.DiscountPercentage = sales.DiscountPercentage
	m.CustomerId = sales.Customer.Id

	if sales.CustomerPay < 1 {
		m.PaymentStatus = "pending"
	} else if sales.CustomerPay < sales.Total {
		m.PaymentStatus = "down_payment"
	} else {
		m.PaymentStatus = "complete"
	}

	if sales.SaleStatus == "" {
		m.SaleStatus = "complete"
	}

	db.Create(&m)
	salesRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(m.Id))
	return salesRes, nil
}

func (repository SalesRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, sales *domain.Sale) (res domain.Sale, err error) {
	if err := db.Where("id = ?", sales.Id).First(&res).Error; err != nil {
		return res, errors.New("NOT_FOUND|penjualan tidak ditemukan")
	}

	db.Model(&res).Updates(&sales)

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
	db.Order("sales.id desc").Find(&salesRes)
	return salesRes, nil
}

func (repository SalesRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.SaleDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("sales").
		Select(`
			sales.*,
			stores.id store_id, stores.name store_name, 
			customers.id customer_id, customers.name customer_name, 
			cash_registers.id cash_register_id, cash_registers.cash_in_hand cash_register_cash_in_hand,
			users.name created_by_name
		`).
		Joins(`
			left join stores on stores.id = sales.store_id
			left join customers on customers.id = sales.customer_id
			left join cash_registers on cash_registers.id = sales.cash_register_id
			JOIN users ON users.id = sales.created_by
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

func (repository SalesRepositoryImpl) ReportDatatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string, filter map[string]string) (datatableRes []web.SaleDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("sales")
	qry.Select(`
			sales.*,
			stores.id store_id, stores.name store_name, 
			customers.id customer_id, customers.name customer_name, 
			cash_registers.id cash_register_id, cash_registers.cash_in_hand cash_register_cash_in_hand ,
			users.name created_by_name
		`)
	qry.Joins(`
			left join stores on stores.id = sales.store_id
			left join customers on customers.id = sales.customer_id
			left join cash_registers on cash_registers.id = sales.cash_register_id
			JOIN users ON users.id = sales.created_by
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

func (repository SalesRepositoryImpl) FindByCreatedAt(ctx echo.Context, db *gorm.DB, dateRange *web.DateRange) (saleRes []web.SalesPosGet, err error) {
	qry := db.Table("sales")
	qry.Select(`
			sales.*,
			stores.id store_id, stores.name store_name, 
			customers.id customer_id, customers.name customer_name, 
			cash_registers.id cash_register_id, cash_registers.cash_in_hand cash_in_hand,
			users.name created_by_name
		`)
	qry.Joins(`
			left join stores on stores.id = sales.store_id
			left join customers on customers.id = sales.customer_id
			left join cash_registers on cash_registers.id = sales.cash_register_id
			JOIN users ON users.id = sales.created_by
		`)
	if dateRange.StartDate != "" && dateRange.EndDate != ""{
		qry.Where("(sales.created_at > ? AND sales.created_at < ?)", dateRange.StartDate, dateRange.EndDate)
	}
	qry.Find(&saleRes)
	return saleRes, nil
}