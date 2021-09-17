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
	ExpenseService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	ExpenseServiceImpl struct {
		ExpenseRepository repository.ExpenseRepository
		db *gorm.DB
	}
)

func NewExpenseService(ExpenseRepository repository.ExpenseRepository, db *gorm.DB) ExpenseService {
	return &ExpenseServiceImpl{
		ExpenseRepository: ExpenseRepository,
		db: db,
	}
}

func (service *ExpenseServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.Expense)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	expenseRepo, err := service.ExpenseRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", expenseRepo), err
}

func (service ExpenseServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Expense)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	expenseRepo, err := service.ExpenseRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", expenseRepo), err
}

func (service ExpenseServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Expense)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.ExpenseRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service ExpenseServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	expenseRepo, err := service.ExpenseRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", expenseRepo), err
}

func (service ExpenseServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	expenseRepo, err := service.ExpenseRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", expenseRepo), err
}

