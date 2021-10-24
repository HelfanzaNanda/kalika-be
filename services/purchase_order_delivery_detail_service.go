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
	PurchaseOrderDeliveryDetailService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	PurchaseOrderDeliveryDetailServiceImpl struct {
		PurchaseOrderDeliveryDetailRepository repository.PurchaseOrderDeliveryDetailRepository
		db *gorm.DB
	}
)

func NewPurchaseOrderDeliveryDetailService(PurchaseOrderDeliveryDetailRepository repository.PurchaseOrderDeliveryDetailRepository, db *gorm.DB) PurchaseOrderDeliveryDetailService {
	return &PurchaseOrderDeliveryDetailServiceImpl{
		PurchaseOrderDeliveryDetailRepository: PurchaseOrderDeliveryDetailRepository,
		db: db,
	}
}

func (service *PurchaseOrderDeliveryDetailServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new([]domain.PurchaseOrderDeliveryDetail)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	purchaseOrderDeliveryDetailRepo, err := service.PurchaseOrderDeliveryDetailRepository.Create(ctx, tx, *o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", purchaseOrderDeliveryDetailRepo), err
}

func (service PurchaseOrderDeliveryDetailServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.PurchaseOrderDeliveryDetail)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	purchaseOrderDeliveryDetailRepo, err := service.PurchaseOrderDeliveryDetailRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", purchaseOrderDeliveryDetailRepo), err
}

func (service PurchaseOrderDeliveryDetailServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.PurchaseOrderDeliveryDetail)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.PurchaseOrderDeliveryDetailRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service PurchaseOrderDeliveryDetailServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	purchaseOrderDeliveryDetailRepo, err := service.PurchaseOrderDeliveryDetailRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", purchaseOrderDeliveryDetailRepo), err
}

func (service PurchaseOrderDeliveryDetailServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	purchaseOrderDeliveryDetailRepo, err := service.PurchaseOrderDeliveryDetailRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", purchaseOrderDeliveryDetailRepo), err
}

