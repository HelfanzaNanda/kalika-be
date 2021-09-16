package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	RawMaterialRepository interface{
		Create(ctx echo.Context, db *gorm.DB, rawMaterial *domain.RawMaterial) (domain.RawMaterial, error)
		Update(ctx echo.Context, db *gorm.DB, rawMaterial *domain.RawMaterial) (domain.RawMaterial, error)
		Delete(ctx echo.Context, db *gorm.DB, rawMaterial *domain.RawMaterial) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.RawMaterial, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.RawMaterial, error)
	}

	RawMaterialRepositoryImpl struct {

	}
)

func NewRawMaterialRepository() RawMaterialRepository {
	return &RawMaterialRepositoryImpl{}
}

func (repository RawMaterialRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, rawMaterial *domain.RawMaterial) (domain.RawMaterial, error) {
	db.Create(&rawMaterial)
	rawMaterialRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(rawMaterial.Id))
	return rawMaterialRes, nil
}

func (repository RawMaterialRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, rawMaterial *domain.RawMaterial) (domain.RawMaterial, error) {
	db.Where("id = ?", rawMaterial.Id).Updates(&rawMaterial)
	rawMaterialRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(rawMaterial.Id))
	return rawMaterialRes, nil
}

func (repository RawMaterialRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, rawMaterial *domain.RawMaterial) (bool, error) {
	results := db.Where("id = ?", rawMaterial.Id).Delete(&rawMaterial)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|rawMaterial tidak ditemukan")
	}
	return true, nil
}

func (repository RawMaterialRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (rawMaterialRes domain.RawMaterial, err error) {
	results := db.Where(key+" = ?", value).First(&rawMaterialRes)
	if results.RowsAffected < 1 {
		return rawMaterialRes, errors.New("NOT_FOUND|rawMaterial tidak ditemukan")
	}
	return rawMaterialRes, nil
}

func (repository RawMaterialRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (rawMaterialRes []domain.RawMaterial, err error) {
	db.Find(&rawMaterialRes)
	return rawMaterialRes, nil
}

