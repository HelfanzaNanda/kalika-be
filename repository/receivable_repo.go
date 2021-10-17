package repository

import (
	"errors"
	"fmt"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"
	"time"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type (
	ReceivableRepository interface{
		Create(ctx echo.Context, db *gorm.DB, receivable *web.ReceivablePosPost) (domain.Receivable, error)
		Update(ctx echo.Context, db *gorm.DB, receivable *web.ReceivablePosPost) (domain.Receivable, error)
		Delete(ctx echo.Context, db *gorm.DB, receivable *domain.Receivable) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Receivable, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Receivable, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.ReceivableDatatable, int64, int64, error)
		ReportDatatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string, filter map[string]string) ([]web.ReceivableDatatable, int64, int64, error)
		FindByCreatedAt(ctx echo.Context, db *gorm.DB, dateRange *web.DateRange) ([]web.ReceivableGet, error)
	}

	ReceivableRepositoryImpl struct {

	}
)

func NewReceivableRepository() ReceivableRepository {
	return &ReceivableRepositoryImpl{}
}

func (repository ReceivableRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, receivable *web.ReceivablePosPost) (domain.Receivable, error) {
	layoutFormat := "2006-01-02"
	date, err := time.Parse(layoutFormat, receivable.Date)
	if err != nil{
		fmt.Println("########## ERROR TIME PARSE")
		fmt.Println(err)
		fmt.Println(receivable)
		fmt.Println("########## ERROR TIME PARSE")
	}
	model := domain.Receivable{}
	model.Model = receivable.Model
	model.ModelId = receivable.ModelId
	model.CustomerId = receivable.CustomerId
	model.StoreConsignmentId = receivable.StoreConsignmentId
	model.Receivables = receivable.Receivables
	model.Total = receivable.Total
	model.Note = receivable.Note
	model.Date = date
	model.CreatedBy = helpers.StringToInt(ctx.Get("userInfo").(map[string]interface{})["id"].(string))
	db.Create(&model)
	receivableRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(model.Id))
	return receivableRes, nil
}

func (repository ReceivableRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, receivable *web.ReceivablePosPost) (domain.Receivable, error) {
	layoutFormat := "2006-01-02"
	date, err := time.Parse(layoutFormat, receivable.Date)
	if err != nil{
		fmt.Println("########## ERROR TIME PARSE")
		fmt.Println(err)
		fmt.Println(receivable)
		fmt.Println("########## ERROR TIME PARSE")
	}
	model := domain.Receivable{}
	model.CustomerId = receivable.CustomerId
	model.StoreConsignmentId = receivable.StoreConsignmentId
	model.Receivables = receivable.Receivables
	model.Total = receivable.Total
	model.Note = receivable.Note
	model.Date = date
	model.CreatedBy = helpers.StringToInt(ctx.Get("userInfo").(map[string]interface{})["id"].(string))
	db.Where("id = ?", receivable.Id).Updates(&model)
	receivableRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(receivable.Id))
	return receivableRes, nil
}

func (repository ReceivableRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, receivable *domain.Receivable) (bool, error) {
	results := db.Where("id = ?", receivable.Id).Delete(&receivable)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|receivable tidak ditemukan")
	}
	return true, nil
}

func (repository ReceivableRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (receivableRes domain.Receivable, err error) {
	results := db.Where(key+" = ?", value).First(&receivableRes)
	if results.RowsAffected < 1 {
		return receivableRes, errors.New("NOT_FOUND|receivable tidak ditemukan")
	}
	return receivableRes, nil
}

func (repository ReceivableRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (receivableRes []domain.Receivable, err error) {
	db.Find(&receivableRes)
	return receivableRes, nil
}


func (repository ReceivableRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.ReceivableDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("receivables")
	qry.Select(`
		receivables.*,
		users.id user_id, users.name user_name,
		customers.id customer_id, customers.name customer_name,
		store_consignments.id store_consignment_id, store_consignments.store_name store_consignment_name
	`)
	qry.Joins(`
		left join users on users.id = receivables.created_by
		left join customers on customers.id = receivables.customer_id
		left join store_consignments on store_consignments.id = receivables.store_consignment_id
	`)
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(receivables.id = ? OR receivables.date LIKE ?)", search, "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("receivables.id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}


func (repository ReceivableRepositoryImpl) ReportDatatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string, filter map[string]string) (datatableRes []web.ReceivableDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("receivables")
	qry.Select(`
		receivables.*,
		users.id user_id, users.name user_name,
		customers.id customer_id, customers.name customer_name,
		store_consignments.id store_consignment_id, store_consignments.store_name store_consignment_name
	`)
	qry.Joins(`
		left join users on users.id = receivables.created_by
		left join customers on customers.id = receivables.customer_id
		left join store_consignments on store_consignments.id = receivables.store_consignment_id
	`)
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(receivables.id = ? OR receivables.date LIKE ?)", search, "%"+search+"%")
	}
	if filter["start_date"] != "" && filter["end_date"] != "" {
		qry.Where("(receivables.created_at > ? AND receivables.created_at < ?)", filter["start_date"], filter["end_date"])
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("receivables.id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}

func (repository ReceivableRepositoryImpl) FindByCreatedAt(ctx echo.Context, db *gorm.DB, dateRange *web.DateRange) (receivableRes []web.ReceivableGet, err error) {
	qry := db.Table("receivables")
	qry.Select(`
		receivables.*,
		users.id user_id, users.name created_by_name,
		customers.id customer_id, customers.name customer_name,
		store_consignments.id store_consignment_id, store_consignments.store_name store_consignment_name
	`)
	qry.Joins(`
		join users on users.id = receivables.created_by
		left join customers on customers.id = receivables.customer_id
		left join store_consignments on store_consignments.id = receivables.store_consignment_id
	`)
	if dateRange.StartDate != "" && dateRange.EndDate != ""{
		qry.Where("(receivables.created_at > ? AND receivables.created_at < ?)", dateRange.StartDate, dateRange.EndDate)
	}
	qry.Order("receivables.id desc")
	qry.Find(&receivableRes)
	return receivableRes, nil
}