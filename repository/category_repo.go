package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	CategoryRepository interface{
		Create(ctx echo.Context, db *gorm.DB, category *domain.Category) (domain.Category, error)
		Update(ctx echo.Context, db *gorm.DB, category *domain.Category) (domain.Category, error)
		Delete(ctx echo.Context, db *gorm.DB, category *domain.Category) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Category, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Category, error)
	}

	CategoryRepositoryImpl struct {

	}
)

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository CategoryRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, category *domain.Category) (domain.Category, error) {
	db.Create(&category)
	categoryRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(category.Id))
	return categoryRes, nil
}

func (repository CategoryRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, category *domain.Category) (domain.Category, error) {
	db.Where("id = ?", category.Id).Updates(&category)
	categoryRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(category.Id))
	return categoryRes, nil
}

func (repository CategoryRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, category *domain.Category) (bool, error) {
	results := db.Where("id = ?", category.Id).Delete(&category)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|category tidak ditemukan")
	}
	return true, nil
}

func (repository CategoryRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (categoryRes domain.Category, err error) {
	results := db.Where(key+" = ?", value).First(&categoryRes)
	if results.RowsAffected < 1 {
		return categoryRes, errors.New("NOT_FOUND|category tidak ditemukan")
	}
	return categoryRes, nil
}

func (repository CategoryRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (categoryRes []domain.Category, err error) {
	db.Find(&categoryRes)
	return categoryRes, nil
}

