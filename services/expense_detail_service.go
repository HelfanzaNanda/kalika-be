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
	ExpenseDetailService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	ExpenseDetailServiceImpl struct {
		ExpenseDetailRepository repository.ExpenseDetailRepository
		db *gorm.DB
	}
)

func NewExpenseDetailService(ExpenseDetailRepository repository.ExpenseDetailRepository, db *gorm.DB) ExpenseDetailService {
	return &ExpenseDetailServiceImpl{
		ExpenseDetailRepository: ExpenseDetailRepository,
		db: db,
	}
}

func (service *ExpenseDetailServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(web.ExpensePosPost)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	expenseDetailRepo, err := service.ExpenseDetailRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", expenseDetailRepo), err
}

func (service ExpenseDetailServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.ExpenseDetail)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	expenseDetailRepo, err := service.ExpenseDetailRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", expenseDetailRepo), err
}

func (service ExpenseDetailServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.ExpenseDetail)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.ExpenseDetailRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service ExpenseDetailServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	expenseDetailRepo, err := service.ExpenseDetailRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", expenseDetailRepo), err
}

func (service ExpenseDetailServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)
	params := make(map[string][]string)
	expenseDetailRepo, err := service.ExpenseDetailRepository.FindAll(ctx, tx, params)

	return helpers.Response("OK", "Sukses Mengambil Data", expenseDetailRepo), err
}

