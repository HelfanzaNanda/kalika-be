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
	DebtDetailService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	DebtDetailServiceImpl struct {
		DebtDetailRepository repository.DebtDetailRepository
		db *gorm.DB
	}
)

func NewDebtDetailService(debtDetailRepository repository.DebtDetailRepository, db *gorm.DB) DebtDetailService {
	return &DebtDetailServiceImpl{
		DebtDetailRepository: debtDetailRepository,
		db: db,
	}
}

func (service *DebtDetailServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.DebtDetail)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	debtDetailRepo, err := service.DebtDetailRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", debtDetailRepo), err
}

func (service DebtDetailServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.DebtDetail)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	debtDetailRepo, err := service.DebtDetailRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", debtDetailRepo), err
}

func (service DebtDetailServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.DebtDetail)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.DebtDetailRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service DebtDetailServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	debtDetailRepo, err := service.DebtDetailRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", debtDetailRepo), err
}

func (service DebtDetailServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	debtDetailRepo, err := service.DebtDetailRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", debtDetailRepo), err
}

