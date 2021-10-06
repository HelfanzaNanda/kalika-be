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
	ExpenseCategoryRepository interface{
		Create(ctx echo.Context, db *gorm.DB, expenseCategory *domain.ExpenseCategory) (domain.ExpenseCategory, error)
		Update(ctx echo.Context, db *gorm.DB, expenseCategory *domain.ExpenseCategory) (domain.ExpenseCategory, error)
		Delete(ctx echo.Context, db *gorm.DB, expenseCategory *domain.ExpenseCategory) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.ExpenseCategory, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.ExpenseCategory, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.ExpenseCategoryDatatable, int64, int64, error)
	}

	ExpenseCategoryRepositoryImpl struct {

	}
)

func NewExpenseCategoryRepository() ExpenseCategoryRepository {
	return &ExpenseCategoryRepositoryImpl{}
}

func (repository ExpenseCategoryRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, expenseCategory *domain.ExpenseCategory) (domain.ExpenseCategory, error) {
	db.Create(&expenseCategory)
	expenseCategoryRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(expenseCategory.Id))
	return expenseCategoryRes, nil
}

func (repository ExpenseCategoryRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, expenseCategory *domain.ExpenseCategory) (domain.ExpenseCategory, error) {
	db.Where("id = ?", expenseCategory.Id).Updates(&expenseCategory)
	expenseCategoryRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(expenseCategory.Id))
	return expenseCategoryRes, nil
}

func (repository ExpenseCategoryRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, expenseCategory *domain.ExpenseCategory) (bool, error) {
	results := db.Where("id = ?", expenseCategory.Id).Delete(&expenseCategory)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|expenseCategory tidak ditemukan")
	}
	return true, nil
}

func (repository ExpenseCategoryRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (expenseCategoryRes domain.ExpenseCategory, err error) {
	results := db.Where(key+" = ?", value).First(&expenseCategoryRes)
	if results.RowsAffected < 1 {
		return expenseCategoryRes, errors.New("NOT_FOUND|expenseCategory tidak ditemukan")
	}
	return expenseCategoryRes, nil
}

func (repository ExpenseCategoryRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (expenseCategoryRes []domain.ExpenseCategory, err error) {
	db.Find(&expenseCategoryRes)
	return expenseCategoryRes, nil
}

func (repository ExpenseCategoryRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.ExpenseCategoryDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("expense_categories")

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