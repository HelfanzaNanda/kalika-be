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
	SalesReturnDetailService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	SalesReturnDetailServiceImpl struct {
		SalesReturnDetailRepository repository.SalesReturnDetailRepository
		db *gorm.DB
	}
)

func NewSalesReturnDetailService(SalesReturnDetailRepository repository.SalesReturnDetailRepository, db *gorm.DB) SalesReturnDetailService {
	return &SalesReturnDetailServiceImpl{
		SalesReturnDetailRepository: SalesReturnDetailRepository,
		db: db,
	}
}

func (service *SalesReturnDetailServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(web.SalesReturnPost)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	salesReturnDetailRepo, err := service.SalesReturnDetailRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", salesReturnDetailRepo), err
}

func (service SalesReturnDetailServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.SalesReturnDetail)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	salesReturnDetailRepo, err := service.SalesReturnDetailRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", salesReturnDetailRepo), err
}

func (service SalesReturnDetailServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.SalesReturnDetail)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.SalesReturnDetailRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service SalesReturnDetailServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	salesReturnDetailRepo, err := service.SalesReturnDetailRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", salesReturnDetailRepo), err
}

func (service SalesReturnDetailServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	salesReturnDetailRepo, err := service.SalesReturnDetailRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", salesReturnDetailRepo), err
}

