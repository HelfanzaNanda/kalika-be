package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	DebtRepository interface{
		Create(ctx echo.Context, db *gorm.DB, debt *domain.Debt) (domain.Debt, error)
		Update(ctx echo.Context, db *gorm.DB, debt *domain.Debt) (domain.Debt, error)
		Delete(ctx echo.Context, db *gorm.DB, debt *domain.Debt) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Debt, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Debt, error)
	}

	DebtRepositoryImpl struct {

	}
)

func NewDebtRepository() DebtRepository {
	return &DebtRepositoryImpl{}
}

func (repository DebtRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, debt *domain.Debt) (domain.Debt, error) {
	db.Create(&debt)
	debtRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(debt.Id))
	return debtRes, nil
}

func (repository DebtRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, debt *domain.Debt) (domain.Debt, error) {
	db.Where("id = ?", debt.Id).Updates(&debt)
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

