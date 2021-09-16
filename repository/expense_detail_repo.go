package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	ExpenseDetailRepository interface{
		Create(ctx echo.Context, db *gorm.DB, expenseDetail *domain.ExpenseDetail) (domain.ExpenseDetail, error)
		Update(ctx echo.Context, db *gorm.DB, expenseDetail *domain.ExpenseDetail) (domain.ExpenseDetail, error)
		Delete(ctx echo.Context, db *gorm.DB, expenseDetail *domain.ExpenseDetail) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.ExpenseDetail, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.ExpenseDetail, error)
	}

	ExpenseDetailRepositoryImpl struct {

	}
)

func NewExpenseDetailRepository() ExpenseDetailRepository {
	return &ExpenseDetailRepositoryImpl{}
}

func (repository ExpenseDetailRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, expenseDetail *domain.ExpenseDetail) (domain.ExpenseDetail, error) {
	db.Create(&expenseDetail)
	expenseDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(expenseDetail.Id))
	return expenseDetailRes, nil
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

func (repository ExpenseDetailRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (expenseDetailRes domain.ExpenseDetail, err error) {
	results := db.Where(key+" = ?", value).First(&expenseDetailRes)
	if results.RowsAffected < 1 {
		return expenseDetailRes, errors.New("NOT_FOUND|expenseDetail tidak ditemukan")
	}
	return expenseDetailRes, nil
}

func (repository ExpenseDetailRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (expenseDetailRes []domain.ExpenseDetail, err error) {
	db.Find(&expenseDetailRes)
	return expenseDetailRes, nil
}

