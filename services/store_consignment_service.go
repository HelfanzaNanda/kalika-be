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
	StoreConsignmentService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
		Datatable(ctx echo.Context) (res web.Datatable, err error)

	}

	StoreConsignmentServiceImpl struct {
		StoreConsignmentRepository repository.StoreConsignmentRepository
		db *gorm.DB
	}
)

func NewStoreConsignmentService(StoreConsignmentRepository repository.StoreConsignmentRepository, db *gorm.DB) StoreConsignmentService {
	return &StoreConsignmentServiceImpl{
		StoreConsignmentRepository: StoreConsignmentRepository,
		db: db,
	}
}

func (service *StoreConsignmentServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	storeConsignmentRepo := domain.StoreConsignment{}
	o := new(domain.StoreConsignment)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if o.Id > 0 {
		storeConsignmentRepo, err = service.StoreConsignmentRepository.Update(ctx, tx, o)
	} else {
		storeConsignmentRepo, err = service.StoreConsignmentRepository.Create(ctx, tx, o)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", storeConsignmentRepo), err
}

func (service StoreConsignmentServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.StoreConsignment)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	storeConsignmentRepo, err := service.StoreConsignmentRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", storeConsignmentRepo), err
}

func (service StoreConsignmentServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.StoreConsignment)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.StoreConsignmentRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service StoreConsignmentServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	storeConsignmentRepo, err := service.StoreConsignmentRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", storeConsignmentRepo), err
}

func (service StoreConsignmentServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	storeConsignmentRepo, err := service.StoreConsignmentRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", storeConsignmentRepo), err
}


func (service *StoreConsignmentServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params,_ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))

	storeConsignmentRepo, totalData, totalFiltered, _ := service.StoreConsignmentRepository.Datatable(ctx, tx, draw, limit, start, search)

	data := make([]interface{}, 0)
	for _, v := range storeConsignmentRepo {
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