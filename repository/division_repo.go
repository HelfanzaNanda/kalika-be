package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	DivisionRepository interface{
		Create(ctx echo.Context, db *gorm.DB, division *domain.Division) (domain.Division, error)
		Update(ctx echo.Context, db *gorm.DB, division *domain.Division) (domain.Division, error)
		Delete(ctx echo.Context, db *gorm.DB, division *domain.Division) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Division, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Division, error)
	}

	DivisionRepositoryImpl struct {

	}
)

func NewDivisionRepository() DivisionRepository {
	return &DivisionRepositoryImpl{}
}

func (d DivisionRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, division *domain.Division) (domain.Division, error) {
	db.Create(&division)
	DivisionRes,_ := d.FindById(ctx, db, "id", helpers.IntToString(division.Id))
	return DivisionRes, nil
}

func (d DivisionRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, division *domain.Division) (domain.Division, error) {
	db.Where("id = ?", division.Id).Updates(&division)
	DivisionRes,_ := d.FindById(ctx, db, "id", helpers.IntToString(division.Id))
	return DivisionRes, nil
}

func (d DivisionRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, division *domain.Division) (bool, error) {
	results := db.Where("id = ?", division.Id).Delete(&division)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|division tidak ditemukan")
	}
	return true, nil
}

func (d DivisionRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (DivisionRes domain.Division, err error) {
	results := db.Where(key+" = ?", value).First(&DivisionRes)
	if results.RowsAffected < 1 {
		return DivisionRes, errors.New("NOT_FOUND|division tidak ditemukan")
	}
	return DivisionRes, nil
}

func (d DivisionRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (res []domain.Division, err error) {
	db.Find(&res)
	return res, nil
}

