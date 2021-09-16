package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	CustomOrderRepository interface{
		Create(ctx echo.Context, db *gorm.DB, customOrder *domain.CustomOrder) (domain.CustomOrder, error)
		Update(ctx echo.Context, db *gorm.DB, customOrder *domain.CustomOrder) (domain.CustomOrder, error)
		Delete(ctx echo.Context, db *gorm.DB, customOrder *domain.CustomOrder) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.CustomOrder, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.CustomOrder, error)
	}

	CustomOrderRepositoryImpl struct {

	}
)

func NewCustomOrderRepository() CustomOrderRepository {
	return &CustomOrderRepositoryImpl{}
}

func (repository CustomOrderRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, customOrder *domain.CustomOrder) (domain.CustomOrder, error) {
	db.Create(&customOrder)
	customOrderRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(customOrder.Id))
	return customOrderRes, nil
}

func (repository CustomOrderRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, customOrder *domain.CustomOrder) (domain.CustomOrder, error) {
	db.Where("id = ?", customOrder.Id).Updates(&customOrder)
	customOrderRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(customOrder.Id))
	return customOrderRes, nil
}

func (repository CustomOrderRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, customOrder *domain.CustomOrder) (bool, error) {
	results := db.Where("id = ?", customOrder.Id).Delete(&customOrder)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|customOrder tidak ditemukan")
	}
	return true, nil
}

func (repository CustomOrderRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (customOrderRes domain.CustomOrder, err error) {
	results := db.Where(key+" = ?", value).First(&customOrderRes)
	if results.RowsAffected < 1 {
		return customOrderRes, errors.New("NOT_FOUND|customOrder tidak ditemukan")
	}
	return customOrderRes, nil
}

func (repository CustomOrderRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (customOrderRes []domain.CustomOrder, err error) {
	db.Find(&customOrderRes)
	return customOrderRes, nil
}

