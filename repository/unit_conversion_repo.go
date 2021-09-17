package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	UnitConversionRepository interface{
		Create(ctx echo.Context, db *gorm.DB, unitConversion *domain.UnitConversion) (domain.UnitConversion, error)
		Update(ctx echo.Context, db *gorm.DB, unitConversion *domain.UnitConversion) (domain.UnitConversion, error)
		Delete(ctx echo.Context, db *gorm.DB, unitConversion *domain.UnitConversion) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.UnitConversion, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.UnitConversion, error)
	}

	UnitConversionRepositoryImpl struct {

	}
)

func NewUnitConversionRepository() UnitConversionRepository {
	return &UnitConversionRepositoryImpl{}
}

func (repository UnitConversionRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, unitConversion *domain.UnitConversion) (domain.UnitConversion, error) {
	db.Create(&unitConversion)
	unitConversionRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(unitConversion.Id))
	return unitConversionRes, nil
}

func (repository UnitConversionRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, unitConversion *domain.UnitConversion) (domain.UnitConversion, error) {
	db.Where("id = ?", unitConversion.Id).Updates(&unitConversion)
	unitConversionRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(unitConversion.Id))
	return unitConversionRes, nil
}

func (repository UnitConversionRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, unitConversion *domain.UnitConversion) (bool, error) {
	results := db.Where("id = ?", unitConversion.Id).Delete(&unitConversion)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|unitConversion tidak ditemukan")
	}
	return true, nil
}

func (repository UnitConversionRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (unitConversionRes domain.UnitConversion, err error) {
	results := db.Where(key+" = ?", value).First(&unitConversionRes)
	if results.RowsAffected < 1 {
		return unitConversionRes, errors.New("NOT_FOUND|unitConversion tidak ditemukan")
	}
	return unitConversionRes, nil
}

func (repository UnitConversionRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (unitConversionRes []domain.UnitConversion, err error) {
	db.Find(&unitConversionRes)
	return unitConversionRes, nil
}

