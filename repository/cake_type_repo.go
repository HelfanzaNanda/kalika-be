package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	CakeTypeRepository interface{
		Create(ctx echo.Context, db *gorm.DB, cakeType *domain.CakeType) (domain.CakeType, error)
		Update(ctx echo.Context, db *gorm.DB, cakeType *domain.CakeType) (domain.CakeType, error)
		Delete(ctx echo.Context, db *gorm.DB, cakeType *domain.CakeType) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.CakeType, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.CakeType, error)
	}

	CakeTypeRepositoryImpl struct {

	}
)

func NewCakeTypeRepository() CakeTypeRepository {
	return &CakeTypeRepositoryImpl{}
}

func (repository CakeTypeRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, cakeType *domain.CakeType) (domain.CakeType, error) {
	db.Create(&cakeType)
	cakeTypeRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(cakeType.Id))
	return cakeTypeRes, nil
}

func (repository CakeTypeRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, cakeType *domain.CakeType) (domain.CakeType, error) {
	db.Where("id = ?", cakeType.Id).Updates(&cakeType)
	cakeTypeRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(cakeType.Id))
	return cakeTypeRes, nil
}

func (repository CakeTypeRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, cakeType *domain.CakeType) (bool, error) {
	results := db.Where("id = ?", cakeType.Id).Delete(&cakeType)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|cakeType tidak ditemukan")
	}
	return true, nil
}

func (repository CakeTypeRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (cakeTypeRes domain.CakeType, err error) {
	results := db.Where(key+" = ?", value).First(&cakeTypeRes)
	if results.RowsAffected < 1 {
		return cakeTypeRes, errors.New("NOT_FOUND|cakeType tidak ditemukan")
	}
	return cakeTypeRes, nil
}

func (repository CakeTypeRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (cakeTypeRes []domain.CakeType, err error) {
	db.Find(&cakeTypeRes)
	return cakeTypeRes, nil
}

