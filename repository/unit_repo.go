package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	UnitRepository interface{
		Create(ctx echo.Context, db *gorm.DB, unit *domain.Unit) (domain.Unit, error)
		Update(ctx echo.Context, db *gorm.DB, unit *domain.Unit) (domain.Unit, error)
		Delete(ctx echo.Context, db *gorm.DB, unit *domain.Unit) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Unit, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Unit, error)
	}

	UnitRepositoryImpl struct {

	}
)

func NewUnitRepository() UnitRepository {
	return &UnitRepositoryImpl{}
}

func (repository UnitRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, unit *domain.Unit) (domain.Unit, error) {
	db.Create(&unit)
	unitRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(unit.Id))
	return unitRes, nil
}

func (repository UnitRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, unit *domain.Unit) (domain.Unit, error) {
	db.Where("id = ?", unit.Id).Updates(&unit)
	unitRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(unit.Id))
	return unitRes, nil
}

func (repository UnitRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, unit *domain.Unit) (bool, error) {
	results := db.Where("id = ?", unit.Id).Delete(&unit)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|unit tidak ditemukan")
	}
	return true, nil
}

func (repository UnitRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (unitRes domain.Unit, err error) {
	results := db.Where(key+" = ?", value).First(&unitRes)
	if results.RowsAffected < 1 {
		return unitRes, errors.New("NOT_FOUND|unit tidak ditemukan")
	}
	return unitRes, nil
}

func (repository UnitRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (unitRes []domain.Unit, err error) {
	db.Find(&unitRes)
	return unitRes, nil
}

