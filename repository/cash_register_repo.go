package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	CashRegisterRepository interface{
		Create(ctx echo.Context, db *gorm.DB, cashRegister *domain.CashRegister) (domain.CashRegister, error)
		Update(ctx echo.Context, db *gorm.DB, cashRegister *domain.CashRegister) (domain.CashRegister, error)
		Delete(ctx echo.Context, db *gorm.DB, cashRegister *domain.CashRegister) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.CashRegister, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.CashRegister, error)
	}

	CashRegisterRepositoryImpl struct {

	}
)

func NewCashRegisterRepository() CashRegisterRepository {
	return &CashRegisterRepositoryImpl{}
}

func (repository CashRegisterRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, cashRegister *domain.CashRegister) (domain.CashRegister, error) {
	db.Create(&cashRegister)
	cashRegisterRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(cashRegister.Id))
	return cashRegisterRes, nil
}

func (repository CashRegisterRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, cashRegister *domain.CashRegister) (domain.CashRegister, error) {
	db.Where("id = ?", cashRegister.Id).Updates(&cashRegister)
	cashRegisterRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(cashRegister.Id))
	return cashRegisterRes, nil
}

func (repository CashRegisterRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, cashRegister *domain.CashRegister) (bool, error) {
	results := db.Where("id = ?", cashRegister.Id).Delete(&cashRegister)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|cashRegister tidak ditemukan")
	}
	return true, nil
}

func (repository CashRegisterRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (cashRegisterRes domain.CashRegister, err error) {
	results := db.Where(key+" = ?", value).First(&cashRegisterRes)
	if results.RowsAffected < 1 {
		return cashRegisterRes, errors.New("NOT_FOUND|cashRegister tidak ditemukan")
	}
	return cashRegisterRes, nil
}

func (repository CashRegisterRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (cashRegisterRes []domain.CashRegister, err error) {
	db.Find(&cashRegisterRes)
	return cashRegisterRes, nil
}

