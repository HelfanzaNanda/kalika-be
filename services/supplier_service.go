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
	SupplierService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	SupplierServiceImpl struct {
		SupplierRepository repository.SupplierRepository
		db *gorm.DB
	}
)

func NewSupplierService(SupplierRepository repository.SupplierRepository, db *gorm.DB) SupplierService {
	return &SupplierServiceImpl{
		SupplierRepository: SupplierRepository,
		db: db,
	}
}

func (service *SupplierServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.Supplier)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	supplierRepo, err := service.SupplierRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", supplierRepo), err
}

func (service SupplierServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Supplier)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	supplierRepo, err := service.SupplierRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", supplierRepo), err
}

func (service SupplierServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Supplier)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.SupplierRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service SupplierServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	supplierRepo, err := service.SupplierRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", supplierRepo), err
}

func (service SupplierServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	supplierRepo, err := service.SupplierRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", supplierRepo), err
}

