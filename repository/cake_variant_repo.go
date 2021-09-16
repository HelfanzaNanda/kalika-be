package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	CakeVariantRepository interface{
		Create(ctx echo.Context, db *gorm.DB, cakeVariant *domain.CakeVariant) (domain.CakeVariant, error)
		Update(ctx echo.Context, db *gorm.DB, cakeVariant *domain.CakeVariant) (domain.CakeVariant, error)
		Delete(ctx echo.Context, db *gorm.DB, cakeVariant *domain.CakeVariant) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.CakeVariant, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.CakeVariant, error)
	}

	CakeVariantRepositoryImpl struct {

	}
)

func NewCakeVariantRepository() CakeVariantRepository {
	return &CakeVariantRepositoryImpl{}
}

func (repository CakeVariantRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, cakeVariant *domain.CakeVariant) (domain.CakeVariant, error) {
	db.Create(&cakeVariant)
	cakeVariantRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(cakeVariant.Id))
	return cakeVariantRes, nil
}

func (repository CakeVariantRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, cakeVariant *domain.CakeVariant) (domain.CakeVariant, error) {
	db.Where("id = ?", cakeVariant.Id).Updates(&cakeVariant)
	cakeVariantRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(cakeVariant.Id))
	return cakeVariantRes, nil
}

func (repository CakeVariantRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, cakeVariant *domain.CakeVariant) (bool, error) {
	results := db.Where("id = ?", cakeVariant.Id).Delete(&cakeVariant)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|cakeVariant tidak ditemukan")
	}
	return true, nil
}

func (repository CakeVariantRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (cakeVariantRes domain.CakeVariant, err error) {
	results := db.Where(key+" = ?", value).First(&cakeVariantRes)
	if results.RowsAffected < 1 {
		return cakeVariantRes, errors.New("NOT_FOUND|cakeVariant tidak ditemukan")
	}
	return cakeVariantRes, nil
}

func (repository CakeVariantRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (cakeVariantRes []domain.CakeVariant, err error) {
	db.Find(&cakeVariantRes)
	return cakeVariantRes, nil
}

