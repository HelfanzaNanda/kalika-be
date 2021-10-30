package repository

import (
	"errors"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"
	"time"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type (
	SalesReturnRepository interface{
		Create(ctx echo.Context, db *gorm.DB, salesReturn *web.SalesReturnPost) (domain.SalesReturn, error)
		Update(ctx echo.Context, db *gorm.DB, salesReturn *web.SalesReturnPost) (domain.SalesReturn, error)
		Delete(ctx echo.Context, db *gorm.DB, salesReturn *domain.SalesReturn) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.SalesReturn, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]web.SalesReturnGet, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.SalesReturnDatatable, int64, int64, error)
		ReportDatatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string, filter map[string]string) ([]web.SalesReturnDatatable, int64, int64, error)
		FindByCreatedAt(ctx echo.Context, db *gorm.DB, dateRange *web.DateRange) ([]web.SalesReturnPdf, error)
	}

	SalesReturnRepositoryImpl struct {

	}
)

func NewSalesReturnRepository() SalesReturnRepository {
	return &SalesReturnRepositoryImpl{}
}

func (repository SalesReturnRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, salesReturn *web.SalesReturnPost) (domain.SalesReturn, error) {
	storeConsignment := domain.StoreConsignment{}

	if salesReturn.StoreConsignmentId > 0 {
		db.Model(&storeConsignment).Where("id", salesReturn.StoreConsignmentId).First(&storeConsignment)
	}

	model := domain.SalesReturn{}
	model.ModelId = salesReturn.ModelId
	model.Model = salesReturn.Model
	model.Number = "SR"+helpers.IntToString(int(time.Now().Unix()))
	model.CustomerId = salesReturn.CustomerId
	model.StoreConsignmentId = salesReturn.StoreConsignmentId
	model.CreatedBy = helpers.StringToInt(ctx.Get("userInfo").(map[string]interface{})["id"].(string))
	model.Total = salesReturn.Total
	model.Discount = model.Total * (storeConsignment.Discount/100)
	model.Total = model.Total - model.Discount
	db.Create(&model)

	receivableModel := domain.Receivable{}
	receivableDetailsModel := domain.ReceivableDetail{}

	db.Model(&receivableModel).Where("model = ? AND model_id = ?", model.Model, model.ModelId).Update("receivables", gorm.Expr("receivables - ?", model.Total)).First(&receivableModel)

	receivableDetailsModel.ReceivableId = receivableModel.Id
	receivableDetailsModel.Total = model.Total
	receivableDetailsModel.DatePay = time.Now()
	receivableDetailsModel.PaymentMethodId = 0
	receivableDetailsModel.Note = "Retur Konsinyasi "+model.Number
	db.Create(&receivableDetailsModel)

	salesReturnRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(model.Id))
	return salesReturnRes, nil
}

func (repository SalesReturnRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, salesReturn *web.SalesReturnPost) (domain.SalesReturn, error) {
	model := domain.SalesReturn{}
	model.Number = "SR"+helpers.IntToString(int(time.Now().Unix()))
	model.CustomerId = salesReturn.CustomerId
	model.StoreConsignmentId = salesReturn.StoreConsignmentId
	model.Total = salesReturn.Total
	model.CreatedBy = helpers.StringToInt(ctx.Get("userInfo").(map[string]interface{})["id"].(string))
	db.Where("id = ?", salesReturn.Id).Updates(&model)
	
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

func (repository SalesReturnRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (salesReturnRes []web.SalesReturnGet, err error) {
	qry := db.Table("sales_returns").Select("sales_returns.*")
	for k, v := range ctx.QueryParams() {
		if v[0] != "" {
			qry = qry.Where(k+" = ?", v[0])
		}
	}
	qry.Scan(&salesReturnRes)

	for key, val := range salesReturnRes {
		db.Table("sales_return_details").Where("sales_return_id = ?", val.Id).Scan(&salesReturnRes[key].SalesReturnDetail)
	}
	return salesReturnRes, nil
}

func (repository SalesReturnRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.SalesReturnDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("sales_returns").
		Select(`
			sales_returns.*,
			store_consignments.id store_consignment_id, store_consignments.store_name store_consignment_name, 
			customers.id customer_id, customers.name customer_name,
			users.name created_by_name
		`).
		Joins(`
			left join store_consignments on store_consignments.id = sales_returns.store_consignment_id
			left join customers on customers.id = sales_returns.customer_id
			left join users on users.id = sales_returns.created_by
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

func (repository SalesReturnRepositoryImpl) ReportDatatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string, filter map[string]string) (datatableRes []web.SalesReturnDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("sales_returns").
		Select(`
			sales_returns.*,
			store_consignments.id store_consignment_id, store_consignments.store_name store_consignment_name, 
			customers.id customer_id, customers.name customer_name,
			users.name created_by_name
		`).
		Joins(`
			left join store_consignments on store_consignments.id = sales_returns.store_consignment_id
			left join customers on customers.id = sales_returns.customer_id
			left join users on users.id = sales_returns.created_by
		`)

	qry.Count(&totalData)
	if search != "" {
		qry.Where("(sales_returns.id = ? OR sales_returns.number LIKE ?)", search, "%"+search+"%")
	}
	if filter["start_date"] != "" && filter["end_date"] != "" {
		qry.Where("(DATE(sales_returns.created_at) BETWEEN ? AND ?)", filter["start_date"], filter["end_date"])
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("sales_returns.id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}

func (repository SalesReturnRepositoryImpl) FindByCreatedAt(ctx echo.Context, db *gorm.DB, dateRange *web.DateRange) (saleReturnRes []web.SalesReturnPdf, err error) {
	qry := db.Table("sales_returns")
	qry.Select(`
		sales_returns.*,
		store_consignments.id store_consignment_id, store_consignments.store_name store_consignment_name, 
		customers.id customer_id, customers.name customer_name,
		users.name created_by_name
	`)
	qry.Joins(`
		left join store_consignments on store_consignments.id = sales_returns.store_consignment_id
		left join customers on customers.id = sales_returns.customer_id
		join users on users.id = sales_returns.created_by
	`)
	if dateRange.StartDate != "" && dateRange.EndDate != "" {
		qry.Where("(DATE(sales_returns.created_at) BETWEEN ? AND ?)", dateRange.StartDate, dateRange.EndDate)
	}
	qry.Order("sales_returns.id desc")
	qry.Find(&saleReturnRes)
	return saleReturnRes, nil
}