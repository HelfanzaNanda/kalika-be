package services

import (
	//"fmt"
	"strings"

	"github.com/labstack/echo"
	"gorm.io/gorm"

	//"kalika-be/config"
	"kalika-be/helpers"
	"kalika-be/models/web"
	"kalika-be/repository"
)

type (
	ProductLocationService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
		Datatable(ctx echo.Context) (res web.Datatable, err error)
	}

	ProductLocationServiceImpl struct {
		ProductLocationRepository repository.ProductLocationRepository
		db *gorm.DB
	}
)

func NewProductLocationService(ProductLocationRepository repository.ProductLocationRepository, db *gorm.DB) ProductLocationService {
	return &ProductLocationServiceImpl{
		ProductLocationRepository: ProductLocationRepository,
		db: db,
	}
}

func (service *ProductLocationServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	productLocationRepo := web.ProductLocationGet{}
	o := new(web.ProductLocationPost)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if o.Id > 0 {
		productLocationRepo, err = service.ProductLocationRepository.Update(ctx, tx, &o.ProductLocation)
	} else {
		productLocationRepo, err = service.ProductLocationRepository.Create(ctx, tx, &o.ProductLocation)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", productLocationRepo), err
}

func (service ProductLocationServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	return helpers.Response("OK", "Sukses Mengubah Data", "Dummy"), err
}

func (service ProductLocationServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	return helpers.Response("OK", "Sukses Menghapus Data", "Dummy"), err
}

func (service ProductLocationServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	productLocationRepo, err := service.ProductLocationRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", productLocationRepo), err
}

func (service ProductLocationServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	productLocationRepo, err := service.ProductLocationRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", productLocationRepo), err
}

func (service *ProductLocationServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params,_ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))

	productLocationRepo, totalData, totalFiltered, _ := service.ProductLocationRepository.Datatable(ctx, tx, draw, limit, start, search)

	data := make([]interface{}, 0)
	for _, v := range productLocationRepo {
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

