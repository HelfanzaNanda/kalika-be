package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	CustomerRepository interface{
		Create(ctx echo.Context, db *gorm.DB, customer *domain.Customer) (domain.Customer, error)
		Update(ctx echo.Context, db *gorm.DB, customer *domain.Customer) (domain.Customer, error)
		Delete(ctx echo.Context, db *gorm.DB, customer *domain.Customer) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Customer, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Customer, error)
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

