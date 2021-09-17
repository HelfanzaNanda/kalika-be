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
	PurchaseReturnService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	PurchaseReturnServiceImpl struct {
		PurchaseReturnRepository repository.PurchaseReturnRepository
		db *gorm.DB
	}
)

func NewPurchaseReturnService(PurchaseReturnRepository repository.PurchaseReturnRepository, db *gorm.DB) PurchaseReturnService {
	return &PurchaseReturnServiceImpl{
		PurchaseReturnRepository: PurchaseReturnRepository,
		db: db,
	}
}

func (service *PurchaseReturnServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.PurchaseReturn)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	purchaseReturnRepo, err := service.PurchaseReturnRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", purchaseReturnRepo), err
}

func (service PurchaseReturnServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.PurchaseReturn)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	purchaseReturnRepo, err := service.PurchaseReturnRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", purchaseReturnRepo), err
}

func (service PurchaseReturnServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.PurchaseReturn)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.PurchaseReturnRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service PurchaseReturnServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	purchaseReturnRepo, err := service.PurchaseReturnRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", purchaseReturnRepo), err
}

func (service PurchaseReturnServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	purchaseReturnRepo, err := service.PurchaseReturnRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", purchaseReturnRepo), err
}

