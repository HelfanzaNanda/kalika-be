package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	DebtDetailRepository interface{
		Create(ctx echo.Context, db *gorm.DB, debtDetail *domain.DebtDetail) (domain.DebtDetail, error)
		Update(ctx echo.Context, db *gorm.DB, debtDetail *domain.DebtDetail) (domain.DebtDetail, error)
		Delete(ctx echo.Context, db *gorm.DB, debtDetail *domain.DebtDetail) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.DebtDetail, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.DebtDetail, error)
	}

	DebtDetailRepositoryImpl struct {

	}
)

func NewDebtDetailRepository() DebtDetailRepository {
	return &DebtDetailRepositoryImpl{}
}

func (repository DebtDetailRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, debtDetail *domain.DebtDetail) (domain.DebtDetail, error) {
	db.Create(&debtDetail)
	debtDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(debtDetail.Id))
	return debtDetailRes, nil
}

func (repository DebtDetailRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, debtDetail *domain.DebtDetail) (domain.DebtDetail, error) {
	db.Where("id = ?", debtDetail.Id).Updates(&debtDetail)
	debtDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(debtDetail.Id))
	return debtDetailRes, nil
}

func (repository DebtDetailRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, debtDetail *domain.DebtDetail) (bool, error) {
	results := db.Where("id = ?", debtDetail.Id).Delete(&debtDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|debtDetail tidak ditemukan")
	}
	return true, nil
}

func (repository DebtDetailRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (debtDetailRes domain.DebtDetail, err error) {
	results := db.Where(key+" = ?", value).First(&debtDetailRes)
	if results.RowsAffected < 1 {
		return debtDetailRes, errors.New("NOT_FOUND|debtDetail tidak ditemukan")
	}
	return debtDetailRes, nil
}

func (repository DebtDetailRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (debtDetailRes []domain.DebtDetail, err error) {
	db.Find(&debtDetailRes)
	return debtDetailRes, nil
}

