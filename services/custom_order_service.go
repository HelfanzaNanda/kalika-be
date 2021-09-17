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
	CustomOrderService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	CustomOrderServiceImpl struct {
		CustomOrderRepository repository.CustomOrderRepository
		db *gorm.DB
	}
)

func NewCustomOrderService(CustomOrderRepository repository.CustomOrderRepository, db *gorm.DB) CustomOrderService {
	return &CustomOrderServiceImpl{
		CustomOrderRepository: CustomOrderRepository,
		db: db,
	}
}

func (service *CustomOrderServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.CustomOrder)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	customOrderRepo, err := service.CustomOrderRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", customOrderRepo), err
}

func (service CustomOrderServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.CustomOrder)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	customOrderRepo, err := service.CustomOrderRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", customOrderRepo), err
}

func (service CustomOrderServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.CustomOrder)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.CustomOrderRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service CustomOrderServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	customOrderRepo, err := service.CustomOrderRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", customOrderRepo), err
}

func (service CustomOrderServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	customOrderRepo, err := service.CustomOrderRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", customOrderRepo), err
}

