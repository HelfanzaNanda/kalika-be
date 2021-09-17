package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	ExpenseRepository interface{
		Create(ctx echo.Context, db *gorm.DB, expense *domain.Expense) (domain.Expense, error)
		Update(ctx echo.Context, db *gorm.DB, expense *domain.Expense) (domain.Expense, error)
		Delete(ctx echo.Context, db *gorm.DB, expense *domain.Expense) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Expense, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Expense, error)
	}

	ExpenseRepositoryImpl struct {

	}
)

func NewExpenseRepository() ExpenseRepository {
	return &ExpenseRepositoryImpl{}
}

func (repository ExpenseRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, expense *domain.Expense) (domain.Expense, error) {
	db.Create(&expense)
	expenseRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(expense.Id))
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

