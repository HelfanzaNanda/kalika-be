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
	CustomerRepository interface{
		Create(ctx echo.Context, db *gorm.DB, customer *domain.Customer) (domain.Customer, error)
		Update(ctx echo.Context, db *gorm.DB, customer *domain.Customer) (domain.Customer, error)
		Delete(ctx echo.Context, db *gorm.DB, customer *domain.Customer) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Customer, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Customer, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.CustomerDatatable, int64, int64, error)
	}

	CustomerRepositoryImpl struct {

	}
)

func NewCustomerRepository() CustomerRepository {
	return &CustomerRepositoryImpl{}
}

func (repository CustomerRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, customer *domain.Customer) (domain.Customer, error) {
	db.Create(&customer)
	customerRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(customer.Id))
	return customerRes, nil
}

func (repository CustomerRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, customer *domain.Customer) (domain.Customer, error) {
	db.Where("id = ?", customer.Id).Updates(&customer)
	customerRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(customer.Id))
	return customerRes, nil
}

func (repository CustomerRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, customer *domain.Customer) (bool, error) {
	results := db.Where("id = ?", customer.Id).Delete(&customer)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|customer tidak ditemukan")
	}
	return true, nil
}

func (repository CustomerRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (customerRes domain.Customer, err error) {
	results := db.Where(key+" = ?", value).First(&customerRes)
	if results.RowsAffected < 1 {
		return customerRes, errors.New("NOT_FOUND|customer tidak ditemukan")
	}
	return customerRes, nil
}

func (repository CustomerRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (customerRes []domain.Customer, err error) {
	db.Find(&customerRes)
	return customerRes, nil
}

func (repository CustomerRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.CustomerDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("customers")
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