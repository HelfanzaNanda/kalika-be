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
	DebtRepository interface{
		Create(ctx echo.Context, db *gorm.DB, debt *web.DebtPosPost) (domain.Debt, error)
		Update(ctx echo.Context, db *gorm.DB, debt *web.DebtPosPost) (domain.Debt, error)
		Delete(ctx echo.Context, db *gorm.DB, debt *domain.Debt) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Debt, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Debt, error)
		FindByCreatedAt(ctx echo.Context, db *gorm.DB, dateRange *web.DateRange) ([]web.DebtPosGet, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.DebtDatatable, int64, int64, error)
		ReportDatatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string, filter map[string]string) ([]web.DebtDatatable, int64, int64, error)
	}

	DebtRepositoryImpl struct {

	}
)

func NewDebtRepository() DebtRepository {
	return &DebtRepositoryImpl{}
}

func (repository DebtRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, debt *web.DebtPosPost) (domain.Debt, error) {
	layoutFormat := "2006-01-02"
	date, err := time.Parse(layoutFormat, debt.Date)
	if err != nil{
		fmt.Println("########## ERROR TIME PARSE")
		fmt.Println(err)
		fmt.Println(debt)
		fmt.Println("########## ERROR TIME PARSE")
	}
	model := domain.Debt{}
	model.Model = debt.Model
	model.ModelId = debt.ModelId
	model.SupplierId = debt.SupplierId
	model.Debts = debt.Debts
	model.Total = debt.Total
	model.Note = debt.Note
	model.Date = date
	model.CreatedBy = helpers.StringToInt(ctx.Get("userInfo").(map[string]interface{})["id"].(string))
	db.Create(&model)
	debtRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(debt.Id))
	return debtRes, nil
}

func (repository DebtRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, debt *web.DebtPosPost) (domain.Debt, error) {
	layoutFormat := "2006-01-02"
	date, _ := time.Parse(layoutFormat, debt.Date)
	model := domain.Debt{}
	model.SupplierId = debt.SupplierId
	model.Debts = debt.Debts
	model.Total = debt.Total
	model.Note = debt.Note
	model.Date = date
	model.CreatedBy = helpers.StringToInt(ctx.Get("userInfo").(map[string]interface{})["id"].(string))
	db.Where("id = ?", debt.Id).Updates(&model)
	debtRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(debt.Id))
	return debtRes, nil
}

func (repository DebtRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, debt *domain.Debt) (bool, error) {
	results := db.Where("id = ?", debt.Id).Delete(&debt)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|debt tidak ditemukan")
	}
	return true, nil
}

func (repository DebtRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (debtRes domain.Debt, err error) {
	results := db.Where(key+" = ?", value).First(&debtRes)
	if results.RowsAffected < 1 {
		return debtRes, errors.New("NOT_FOUND|debt tidak ditemukan")
	}
	return debtRes, nil
}

func (repository DebtRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (debtRes []domain.Debt, err error) {
	db.Find(&debtRes)
	return debtRes, nil
}


func (repository DebtRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.DebtDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("debts")
	qry.Select(`
		debts.*,
		users.id user_id, users.name user_name,
		suppliers.id supplier_id, suppliers.name supplier_name
	`)
	qry.Joins(`
		left join users on users.id = debts.created_by
		left join suppliers on suppliers.id = debts.supplier_id
	`)
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(debts.id = ? OR debts.date LIKE ?)", search, "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("debts.id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}

func (repository DebtRepositoryImpl) FindByCreatedAt(ctx echo.Context, db *gorm.DB, dateRange *web.DateRange) (debtRes []web.DebtPosGet, err error) {
	qry := db.Table("debts")
	qry.Select(`
		debts.*,
		users.id user_id, users.name created_by_name,
		suppliers.id supplier_id, suppliers.name supplier_name
	`)
	qry.Joins(`
		join users on users.id = debts.created_by
		left join suppliers on suppliers.id = debts.supplier_id
	`)
	if dateRange.StartDate != "" && dateRange.EndDate != ""{
		qry.Where("(DATE(debts.created_at) BETWEEN ? AND ?)", dateRange.StartDate, dateRange.EndDate)
	}
	qry.Order("id desc")
	qry.Find(&debtRes)
	return debtRes, nil
}


func (repository DebtRepositoryImpl) ReportDatatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string, filter map[string]string) (datatableRes []web.DebtDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("debts")
	qry.Select(`
		debts.*,
		users.id user_id, users.name user_name,
		suppliers.id supplier_id, suppliers.name supplier_name
	`)
	qry.Joins(`
		left join users on users.id = debts.created_by
		left join suppliers on suppliers.id = debts.supplier_id
	`)
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(debts.id = ? OR debts.date LIKE ?)", search, "%"+search+"%")
	}
	if filter["start_date"] != "" && filter["end_date"] != "" {
		qry.Where("(DATE(debts.created_at) BETWEEN ? AND ?)", filter["start_date"], filter["end_date"])
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("debts.id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}