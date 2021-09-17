package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	ProductRepository interface{
		Create(ctx echo.Context, db *gorm.DB, product *domain.Product) (domain.Product, error)
		Update(ctx echo.Context, db *gorm.DB, product *domain.Product) (domain.Product, error)
		Delete(ctx echo.Context, db *gorm.DB, product *domain.Product) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Product, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Product, error)
	}

	ProductRepositoryImpl struct {

	}
)

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository ProductRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, product *domain.Product) (domain.Product, error) {
	db.Create(&product)
	productRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(product.Id))
	return productRes, nil
}

func (repository ProductRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, product *domain.Product) (domain.Product, error) {
	db.Where("id = ?", product.Id).Updates(&product)
	productRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(product.Id))
	return productRes, nil
}

func (repository ProductRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, product *domain.Product) (bool, error) {
	results := db.Where("id = ?", product.Id).Delete(&product)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|product tidak ditemukan")
	}
	return true, nil
}

func (repository ProductRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (productRes domain.Product, err error) {
	results := db.Where(key+" = ?", value).First(&productRes)
	if results.RowsAffected < 1 {
		return productRes, errors.New("NOT_FOUND|product tidak ditemukan")
	}
	return productRes, nil
}

func (repository ProductRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (productRes []domain.Product, err error) {
	db.Find(&productRes)
	return productRes, nil
}

