package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	SupplierRepository interface{
		Create(ctx echo.Context, db *gorm.DB, supplier *domain.Supplier) (domain.Supplier, error)
		Update(ctx echo.Context, db *gorm.DB, supplier *domain.Supplier) (domain.Supplier, error)
		Delete(ctx echo.Context, db *gorm.DB, supplier *domain.Supplier) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Supplier, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Supplier, error)
	}

	SupplierRepositoryImpl struct {

	}
)

func NewSupplierRepository() SupplierRepository {
	return &SupplierRepositoryImpl{}
}

func (repository SupplierRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, supplier *domain.Supplier) (domain.Supplier, error) {
	db.Create(&supplier)
	supplierRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(supplier.Id))
	return supplierRes, nil
}

func (repository SupplierRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, supplier *domain.Supplier) (domain.Supplier, error) {
	db.Where("id = ?", supplier.Id).Updates(&supplier)
	supplierRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(supplier.Id))
	return supplierRes, nil
}

func (repository SupplierRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, supplier *domain.Supplier) (bool, error) {
	results := db.Where("id = ?", supplier.Id).Delete(&supplier)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|supplier tidak ditemukan")
	}
	return true, nil
}

func (repository SupplierRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (supplierRes domain.Supplier, err error) {
	results := db.Where(key+" = ?", value).First(&supplierRes)
	if results.RowsAffected < 1 {
		return supplierRes, errors.New("NOT_FOUND|supplier tidak ditemukan")
	}
	return supplierRes, nil
}

func (repository SupplierRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (supplierRes []domain.Supplier, err error) {
	db.Find(&supplierRes)
	return supplierRes, nil
}

