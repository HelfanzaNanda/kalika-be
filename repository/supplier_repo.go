package repository

import (
	"errors"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type (
	SupplierRepository interface{
		Create(ctx echo.Context, db *gorm.DB, supplier *domain.Supplier) (domain.Supplier, error)
		Update(ctx echo.Context, db *gorm.DB, supplier *domain.Supplier) (domain.Supplier, error)
		Delete(ctx echo.Context, db *gorm.DB, supplier *domain.Supplier) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Supplier, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Supplier, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.SupplierDatatable, int64, int64, error)
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

func (repository SupplierRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.SupplierDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("suppliers")
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(id = ? OR name LIKE ?)", search, "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}