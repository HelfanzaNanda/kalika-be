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
	ExpenseDetailRepository interface{
		Create(ctx echo.Context, db *gorm.DB, expenseDetail *web.ExpensePosPost) (web.ExpensePosPost, error)
		Update(ctx echo.Context, db *gorm.DB, expenseDetail *domain.ExpenseDetail) (domain.ExpenseDetail, error)
		Delete(ctx echo.Context, db *gorm.DB, expenseDetail *domain.ExpenseDetail) (bool, error)
		DeleteByExpenseId(ctx echo.Context, db *gorm.DB, expenseId int) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.ExpenseDetail, error)
		FindByExpenseId(ctx echo.Context, db *gorm.DB, expenseId int) ([]domain.ExpenseDetail, error)
		FindAll(ctx echo.Context, db *gorm.DB, params map[string][]string) (web.ExpensePosPost, error)
		ReportDatatable(ctx echo.Context, db *gorm.DB, daterange *web.DateRange) ([]web.ReportExpenseDatatable, error)
	}

	ExpenseDetailRepositoryImpl struct {

	}
)

func NewExpenseDetailRepository() ExpenseDetailRepository {
	return &ExpenseDetailRepositoryImpl{}
}

func (repository ExpenseDetailRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, expenseDetail *web.ExpensePosPost) (res web.ExpensePosPost, err error) {
	var total float64 = 0
	for _, val := range expenseDetail.ExpenseDetails {
		val.ExpenseId = expenseDetail.Id
		db.Create(&val)
		total += val.Amount
		res.ExpenseDetails = append(res.ExpenseDetails, val)
	}
	res.Total = total
	return res, nil
}

func (repository ExpenseDetailRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, expenseDetail *domain.ExpenseDetail) (domain.ExpenseDetail, error) {
	db.Where("id = ?", expenseDetail.Id).Updates(&expenseDetail)
	expenseDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(expenseDetail.Id))
	return expenseDetailRes, nil
}

func (repository ExpenseDetailRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, expenseDetail *domain.ExpenseDetail) (bool, error) {
	results := db.Where("id = ?", expenseDetail.Id).Delete(&expenseDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|expenseDetail tidak ditemukan")
	}
	return true, nil
}
func (repository ExpenseDetailRepositoryImpl) DeleteByExpenseId(ctx echo.Context, db *gorm.DB, expenseId int) (bool, error) {
	results := db.Where("expense_id = ?", expenseId).Delete(domain.ExpenseDetail{})
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|expense Detail tidak ditemukan")
	}
	return true, nil
}

func (repository ExpenseDetailRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (expenseDetailRes domain.ExpenseDetail, err error) {
	results := db.Where(key+" = ?", value).First(&expenseDetailRes)
	if results.RowsAffected < 1 {
		return expenseDetailRes, errors.New("NOT_FOUND|expenseDetail tidak ditemukan")
	}
	return expenseDetailRes, nil
}
func (repository ExpenseDetailRepositoryImpl) FindByExpenseId(ctx echo.Context, db *gorm.DB, expenseId int) (expenseDetailRes []domain.ExpenseDetail, err error) {
	results := db.Where("expense_id = ?", expenseId).Find(&expenseDetailRes)
	if results.RowsAffected < 1 {
		return expenseDetailRes, errors.New("NOT_FOUND|expense Detail tidak ditemukan")
	}
	return expenseDetailRes, nil
}

func (repository ExpenseDetailRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB, params map[string][]string) (expenseDetailRes web.ExpensePosPost, err error) {
	results := db.Table("expense_details").Preload("Expsense")
	for k, v := range params {
		if v[0] != "" && k != "id" {
			results = results.Where(k+" = ?", v[0])
		}
	}

	results.Find(&expenseDetailRes.ExpenseDetails)
	return expenseDetailRes, nil
}

func (repository ExpenseDetailRepositoryImpl) ReportDatatable(ctx echo.Context, db *gorm.DB, daterange *web.DateRange) (datatableRes []web.ReportExpenseDatatable, err error) {
	qry := db.Table("expense_details detail")
	qry.Select("category.name category_name, sum(detail.amount) total ")
	qry.Joins(`
		join expense_categories category on category.id = detail.expense_category_id 
		join expenses expense on expense.id = detail.expense_id
	`)
	if daterange.StartDate != "" && daterange.EndDate != "" {
		qry.Where("(DATE(expense.created_at) BETWEEN ? AND ?)", daterange.StartDate, daterange.EndDate)
	}
	qry.Group("category.name")
	qry.Find(&datatableRes)
	return datatableRes, nil
}
