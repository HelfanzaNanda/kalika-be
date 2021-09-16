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
	PurchaseOrderDeliveryService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	PurchaseOrderDeliveryServiceImpl struct {
		PurchaseOrderDeliveryRepository repository.PurchaseOrderDeliveryRepository
		db *gorm.DB
	}
)

func NewPurchaseOrderDeliveryService(PurchaseOrderDeliveryRepository repository.PurchaseOrderDeliveryRepository, db *gorm.DB) PurchaseOrderDeliveryService {
	return &PurchaseOrderDeliveryServiceImpl{
		PurchaseOrderDeliveryRepository: PurchaseOrderDeliveryRepository,
		db: db,
	}
}

func (service *PurchaseOrderDeliveryServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.PurchaseOrderDelivery)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	purchaseOrderDeliveryRepo, err := service.PurchaseOrderDeliveryRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", purchaseOrderDeliveryRepo), err
}

func (service PurchaseOrderDeliveryServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.PurchaseOrderDelivery)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	purchaseOrderDeliveryRepo, err := service.PurchaseOrderDeliveryRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", purchaseOrderDeliveryRepo), err
}

func (service PurchaseOrderDeliveryServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.PurchaseOrderDelivery)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.PurchaseOrderDeliveryRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service PurchaseOrderDeliveryServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	purchaseOrderDeliveryRepo, err := service.PurchaseOrderDeliveryRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", purchaseOrderDeliveryRepo), err
}

func (service PurchaseOrderDeliveryServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	purchaseOrderDeliveryRepo, err := service.PurchaseOrderDeliveryRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", purchaseOrderDeliveryRepo), err
}

