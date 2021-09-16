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
	SalesReturnService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	SalesReturnServiceImpl struct {
		SalesReturnRepository repository.SalesReturnRepository
		db *gorm.DB
	}
)

func NewSalesReturnService(SalesReturnRepository repository.SalesReturnRepository, db *gorm.DB) SalesReturnService {
	return &SalesReturnServiceImpl{
		SalesReturnRepository: SalesReturnRepository,
		db: db,
	}
}

func (service *SalesReturnServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.SalesReturn)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	salesReturnRepo, err := service.SalesReturnRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", salesReturnRepo), err
}

func (service SalesReturnServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.SalesReturn)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	salesReturnRepo, err := service.SalesReturnRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", salesReturnRepo), err
}

func (service SalesReturnServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.SalesReturn)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.SalesReturnRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service SalesReturnServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	salesReturnRepo, err := service.SalesReturnRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", salesReturnRepo), err
}

func (service SalesReturnServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	salesReturnRepo, err := service.SalesReturnRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", salesReturnRepo), err
}

