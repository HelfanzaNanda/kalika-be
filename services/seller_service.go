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
	SellerService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	SellerServiceImpl struct {
		SellerRepository repository.SellerRepository
		db *gorm.DB
	}
)

func NewSellerService(SellerRepository repository.SellerRepository, db *gorm.DB) SellerService {
	return &SellerServiceImpl{
		SellerRepository: SellerRepository,
		db: db,
	}
}

func (service *SellerServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.Seller)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	sellerRepo, err := service.SellerRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", sellerRepo), err
}

func (service SellerServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Seller)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	sellerRepo, err := service.SellerRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", sellerRepo), err
}

func (service SellerServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Seller)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.SellerRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service SellerServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	sellerRepo, err := service.SellerRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", sellerRepo), err
}

func (service SellerServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	sellerRepo, err := service.SellerRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", sellerRepo), err
}

