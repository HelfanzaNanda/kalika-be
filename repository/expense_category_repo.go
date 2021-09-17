package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	ExpenseCategoryRepository interface{
		Create(ctx echo.Context, db *gorm.DB, expenseCategory *domain.ExpenseCategory) (domain.ExpenseCategory, error)
		Update(ctx echo.Context, db *gorm.DB, expenseCategory *domain.ExpenseCategory) (domain.ExpenseCategory, error)
		Delete(ctx echo.Context, db *gorm.DB, expenseCategory *domain.ExpenseCategory) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.ExpenseCategory, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.ExpenseCategory, error)
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

