package services

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"

	//"kalika-be/config"
	"kalika-be/helpers"
	"kalika-be/models/web"
	"kalika-be/repository"
)

type (
	ProductPriceService interface {
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	ProductPriceServiceImpl struct {
		ProductPriceRepository repository.ProductPriceRepository
		db *gorm.DB
	}
)

func NewProductPriceService(ProductPriceRepository repository.ProductPriceRepository, db *gorm.DB) ProductPriceService {
	return &ProductPriceServiceImpl{
		ProductPriceRepository: ProductPriceRepository,
		db: db,
	}
}

func (service ProductPriceServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	return helpers.Response("OK", "Sukses Mengubah Data", "Dummy"), err
}

func (service ProductPriceServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	return helpers.Response("OK", "Sukses Menghapus Data", "Dummy"), err
}

func (service ProductPriceServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	productPriceRepo, err := service.ProductPriceRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", productPriceRepo), err
}

func (service ProductPriceServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	productPriceRepo, err := service.ProductPriceRepository.FindAll(ctx, tx, map[string][]string{})

	return helpers.Response("OK", "Sukses Mengambil Data", productPriceRepo), err
}