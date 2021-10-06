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
	ExpenseRepository interface{
		Create(ctx echo.Context, db *gorm.DB, expense *web.ExpensePosPost) (domain.Expense, error)
		Update(ctx echo.Context, db *gorm.DB, expense *domain.Expense) (domain.Expense, error)
		Delete(ctx echo.Context, db *gorm.DB, expense *domain.Expense) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Expense, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Expense, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.ExpenseDatatable, int64, int64, error)
	}

	ExpenseRepositoryImpl struct {

	}
)

func NewExpenseRepository() ExpenseRepository {
	return &ExpenseRepositoryImpl{}
}

func (repository ExpenseRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, expense *web.ExpensePosPost) (domain.Expense, error) {
	model := domain.Expense{}
	model.Number = "CS"+helpers.IntToString(int(time.Now().Unix()))
	model.Date = time.Now()
	model.CreatedBy = helpers.StringToInt(ctx.Get("userInfo").(map[string]interface{})["id"].(string))
	db.Create(&model)
	

	expenseRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(model.Id))
	return expenseRes, nil
}

func (repository ExpenseRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, expense *domain.Expense) (domain.Expense, error) {
	db.Where("id = ?", expense.Id).Updates(&expense)
	expenseRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(expense.Id))
	return expenseRes, nil
}

func (repository ExpenseRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, expense *domain.Expense) (bool, error) {
	results := db.Where("id = ?", expense.Id).Delete(&expense)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|expense tidak ditemukan")
	}
	return true, nil
}

func (repository ExpenseRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (expenseRes domain.Expense, err error) {
	results := db.Where(key+" = ?", value).First(&expenseRes)
	if results.RowsAffected < 1 {
		return expenseRes, errors.New("NOT_FOUND|expense tidak ditemukan")
	}
	return expenseRes, nil
}

func (repository ExpenseRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (expenseRes []domain.Expense, err error) {
	db.Find(&expenseRes)
	return expenseRes, nil
}

func (repository ExpenseRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.ExpenseDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("expenses").
	Select("expenses.*, users.name created_by_name").
	Joins("left join users on users.id = expenses.created_by")

	qry.Count(&totalData)
	if search != "" {
		qry.Where("(id = ? OR name LIKE ?)", search, "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}