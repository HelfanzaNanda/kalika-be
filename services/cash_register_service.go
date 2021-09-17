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
	CashRegisterService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	CashRegisterServiceImpl struct {
		CashRegisterRepository repository.CashRegisterRepository
		db *gorm.DB
	}
)

func NewCashRegisterService(CashRegisterRepository repository.CashRegisterRepository, db *gorm.DB) CashRegisterService {
	return &CashRegisterServiceImpl{
		CashRegisterRepository: CashRegisterRepository,
		db: db,
	}
}

func (service *CashRegisterServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.CashRegister)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	cashRegisterRepo, err := service.CashRegisterRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", cashRegisterRepo), err
}

func (service CashRegisterServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.CashRegister)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	cashRegisterRepo, err := service.CashRegisterRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", cashRegisterRepo), err
}

func (service CashRegisterServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.CashRegister)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.CashRegisterRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service CashRegisterServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	cashRegisterRepo, err := service.CashRegisterRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", cashRegisterRepo), err
}

func (service CashRegisterServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	cashRegisterRepo, err := service.CashRegisterRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", cashRegisterRepo), err
}

