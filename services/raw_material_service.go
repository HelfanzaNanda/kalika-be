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
	RawMaterialService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
		Datatable(ctx echo.Context) (res web.Datatable, err error)

	}

	RawMaterialServiceImpl struct {
		RawMaterialRepository repository.RawMaterialRepository
		ProductLocationRepository repository.ProductLocationRepository
		db *gorm.DB
	}
)

func NewRawMaterialService(RawMaterialRepository repository.RawMaterialRepository, productLocationRepository repository.ProductLocationRepository, db *gorm.DB) RawMaterialService {
	return &RawMaterialServiceImpl{
		RawMaterialRepository: RawMaterialRepository,
		ProductLocationRepository: productLocationRepository,
		db: db,
	}
}

func (service *RawMaterialServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	msg := "Sukses Menyimpan Data"
	rawMaterialRepo := domain.RawMaterial{}
	o := new(web.RawMaterialPost)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if o.Id > 0 {
		msg = "Sukses Mengubah Data"
		rawMaterialRepo, err = service.RawMaterialRepository.Update(ctx, tx, &o.RawMaterial)
	} else {
		rawMaterialRepo, err = service.RawMaterialRepository.Create(ctx, tx, &o.RawMaterial)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	for key, _ := range o.ProductLocations {
		o.ProductLocations[key].Model = "RawMaterial"
		o.ProductLocations[key].ProductId = rawMaterialRepo.Id
	}
	if o.Id > 0 {
		service.ProductLocationRepository.DeleteByProduct(ctx, tx, "RawMaterial", o.Id)
		_, err = service.ProductLocationRepository.Create(ctx, tx, o.ProductLocations)
	} else {
		_, err = service.ProductLocationRepository.Create(ctx, tx, o.ProductLocations)
	}
	if err != nil {
		return helpers.Response(err.Error(), "create product location error", nil), err
	}

	return helpers.Response("CREATED", msg, rawMaterialRepo), err
}

func (service RawMaterialServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.RawMaterial)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	rawMaterialRepo, err := service.RawMaterialRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", rawMaterialRepo), err
}

func (service RawMaterialServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.RawMaterial)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.RawMaterialRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service RawMaterialServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	rawMaterialRepo, err := service.RawMaterialRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", rawMaterialRepo), err
}

func (service RawMaterialServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	rawMaterialRepo, err := service.RawMaterialRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", rawMaterialRepo), err
}

func (service *RawMaterialServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params,_ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))

	rawMaterialRepo, totalData, totalFiltered, _ := service.RawMaterialRepository.Datatable(ctx, tx, draw, limit, start, search)
	// if err != nil {
	// 	return helpers.Response(err.Error(), "", nil), err
	// }

	data := make([]interface{}, 0)
	for _, v := range rawMaterialRepo {
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

