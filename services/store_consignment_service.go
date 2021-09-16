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
	StoreConsignmentService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	StoreConsignmentServiceImpl struct {
		StoreConsignmentRepository repository.StoreConsignmentRepository
		db *gorm.DB
	}
)

func NewStoreConsignmentService(StoreConsignmentRepository repository.StoreConsignmentRepository, db *gorm.DB) StoreConsignmentService {
	return &StoreConsignmentServiceImpl{
		StoreConsignmentRepository: StoreConsignmentRepository,
		db: db,
	}
}

func (service *StoreConsignmentServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.StoreConsignment)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	storeConsignmentRepo, err := service.StoreConsignmentRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", storeConsignmentRepo), err
}

func (service StoreConsignmentServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.StoreConsignment)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	storeConsignmentRepo, err := service.StoreConsignmentRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", storeConsignmentRepo), err
}

func (service StoreConsignmentServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.StoreConsignment)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.StoreConsignmentRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service StoreConsignmentServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	storeConsignmentRepo, err := service.StoreConsignmentRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", storeConsignmentRepo), err
}

func (service StoreConsignmentServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	storeConsignmentRepo, err := service.StoreConsignmentRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", storeConsignmentRepo), err
}

