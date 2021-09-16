package services

import (
	//"fmt"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	//"kalika-be/config"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"
	"kalika-be/repository"
)

type (
	CustomerService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	CustomerServiceImpl struct {
		CustomerRepository repository.CustomerRepository
		db *gorm.DB
	}
)

func NewCustomerService(CustomerRepository repository.CustomerRepository, db *gorm.DB) CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository: CustomerRepository,
		db: db,
	}
}

func (service *CustomerServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.Customer)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	customerRepo, err := service.CustomerRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", customerRepo), err
}

func (service CustomerServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Customer)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	customerRepo, err := service.CustomerRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", customerRepo), err
}

func (service CustomerServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Customer)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.CustomerRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service CustomerServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	customerRepo, err := service.CustomerRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", customerRepo), err
}

func (service CustomerServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	customerRepo, err := service.CustomerRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", customerRepo), err
}

