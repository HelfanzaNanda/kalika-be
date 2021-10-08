package services

import (
	//"fmt"
	"strings"

	"github.com/labstack/echo"
	"gorm.io/gorm"

	//"kalika-be/config"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"
	"kalika-be/repository"
)

type (
	ProductService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
		Datatable(ctx echo.Context) (res web.Datatable, err error)

	}

	ProductServiceImpl struct {
		ProductRepository repository.ProductRepository
		ProductPriceRepository repository.ProductPriceRepository
		ProductLocationRepository repository.ProductLocationRepository
		db *gorm.DB
	}
)

func NewProductService(productRepository repository.ProductRepository, productPriceRepository repository.ProductPriceRepository, productLocationRepository repository.ProductLocationRepository, db *gorm.DB) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		ProductPriceRepository: productPriceRepository,
		ProductLocationRepository: productLocationRepository,
		db: db,
	}
}

func (service *ProductServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(web.ProductPosPost)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)
	productRepo, err := service.ProductRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "create product error", nil), err
	}
	o.Id = productRepo.Id
	_, err = service.ProductPriceRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "create product price error", nil), err
	}
	for key, _ := range o.ProductLocations {
		o.ProductLocations[key].Model = "Product"
	}
	_, err = service.ProductLocationRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "create product location error", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", o), err
}

func (service ProductServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(web.ProductPosPost)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)
	_, err = service.ProductRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "update product error", nil), err
	}
	_, err = service.ProductPriceRepository.DeleteByProduct(ctx, tx, id)
	if err != nil {
		return helpers.Response(err.Error(), "delete product price by product error", nil), err
	}

	_, err = service.ProductLocationRepository.DeleteByProduct(ctx, tx, id)
	if err != nil {
		return helpers.Response(err.Error(), "delete product location by product error", nil), err
	}

	o.Id = id
	_, err = service.ProductPriceRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "create many product prices error", nil), err
	}
	for key, _ := range o.ProductLocations {
		o.ProductLocations[key].Model = "Product"
	}
	_, err = service.ProductLocationRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "create many product location error", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", o), err
}

func (service ProductServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Product)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.ProductPriceRepository.DeleteByProduct(ctx, tx, id)
	if err != nil {
		return helpers.Response(err.Error(), "delete prices error ", nil), err
	}
	_, err = service.ProductRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service ProductServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	productPost := web.ProductPosPost{}
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	productRepo, err := service.ProductRepository.FindById(ctx, tx, "id", helpers.IntToString(id))
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}
	productPriceRepo, err := service.ProductPriceRepository.FindByProductId(ctx, tx, id)
	if err != nil {
		productPriceRepo = []domain.ProductPrice{}
	}

	productPost.Product = productRepo
	productPost.ProductPrices = productPriceRepo

	return helpers.Response("OK", "Sukses Mengambil Data", productPost), err
}

func (service ProductServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	productRepo, err := service.ProductRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", productRepo), err
}

func (service *ProductServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params,_ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))

	productRepo, totalData, totalFiltered, _ := service.ProductRepository.Datatable(ctx, tx, draw, limit, start, search)
	// if err != nil {
	// 	return helpers.Response(err.Error(), "", nil), err
	// }

	data := make([]interface{}, 0)
	for _, v := range productRepo {
		v.Action = `<div class="flex">`
		v.Action += `<button type="button" class="btn-edit flex mr-3" id="edit-data" data-id=`+helpers.IntToString(v.Id)+`> <i data-feather="check-square" class="w-4 h-4 mr-1"></i> Edit </button>`
		v.Action += `<button type="button" class="btn-delete flex text-theme-6" id="delete-data" data-id=`+helpers.IntToString(v.Id)+`> <i data-feather="trash-2" class="w-4 h-4 mr-1"></i> Delete </button>`
		v.Action += `</div>`

		data = append(data, v)
	}
	res.Data = data
	res.Order = helpers.ParseFormCollection(ctx.Request(), "order")
	res.Draw = helpers.StringToInt(draw)
	res.RecordsFiltered = totalFiltered
	res.RecordsTotal = totalData

	return res, nil
}

