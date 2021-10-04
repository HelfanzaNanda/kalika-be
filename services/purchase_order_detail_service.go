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
	PurchaseOrderDetailService interface {
		//Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	PurchaseOrderDetailServiceImpl struct {
		PurchaseOrderDetailRepository repository.PurchaseOrderDetailRepository
		db *gorm.DB
	}
)

func NewPurchaseOrderDetailService(PurchaseOrderDetailRepository repository.PurchaseOrderDetailRepository, db *gorm.DB) PurchaseOrderDetailService {
	return &PurchaseOrderDetailServiceImpl{
		PurchaseOrderDetailRepository: PurchaseOrderDetailRepository,
		db: db,
	}
}

//func (service *PurchaseOrderDetailServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
//	o := new(domain.PurchaseOrderDetail)
//	if err := ctx.Bind(o); err != nil {
//		return helpers.Response(err.Error(), "Error Data Binding", nil), err
//	}
//
//	tx := service.db.Begin()
//	defer helpers.CommitOrRollback(tx)
//
//	purchaseOrderDetailRepo, err := service.PurchaseOrderDetailRepository.Create(ctx, tx, o)
//	if err != nil {
//		return helpers.Response(err.Error(), "", nil), err
//	}
//
//	return helpers.Response("CREATED", "Sukses Menyimpan Data", purchaseOrderDetailRepo), err
//}

func (service PurchaseOrderDetailServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.PurchaseOrderDetail)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	purchaseOrderDetailRepo, err := service.PurchaseOrderDetailRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", purchaseOrderDetailRepo), err
}

func (service PurchaseOrderDetailServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.PurchaseOrderDetail)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.PurchaseOrderDetailRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service PurchaseOrderDetailServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	purchaseOrderDetailRepo, err := service.PurchaseOrderDetailRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", purchaseOrderDetailRepo), err
}

func (service PurchaseOrderDetailServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	purchaseOrderDetailRepo, err := service.PurchaseOrderDetailRepository.FindAll(ctx, tx, map[string][]string{})

	return helpers.Response("OK", "Sukses Mengambil Data", purchaseOrderDetailRepo), err
}

