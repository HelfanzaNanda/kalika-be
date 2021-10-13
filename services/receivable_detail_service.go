package services

import (
	//"fmt"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"strings"
	"time"

	//"kalika-be/config"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"
	"kalika-be/repository"
)

type (
	ReceivableDetailService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
		Datatable(ctx echo.Context) (res web.Datatable, err error)
	}

	ReceivableDetailServiceImpl struct {
		ReceivableDetailRepository repository.ReceivableDetailRepository
		db *gorm.DB
	}
)

func NewReceivableDetailService(ReceivableDetailRepository repository.ReceivableDetailRepository, db *gorm.DB) ReceivableDetailService {
	return &ReceivableDetailServiceImpl{
		ReceivableDetailRepository: ReceivableDetailRepository,
		db: db,
	}
}

func (service *ReceivableDetailServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.ReceivableDetail)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	o.DatePay = time.Now()
	receivableDetailRepo, err := service.ReceivableDetailRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", receivableDetailRepo), err
}

func (service *ReceivableDetailServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.ReceivableDetail)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	receivableDetailRepo, err := service.ReceivableDetailRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", receivableDetailRepo), err
}

func (service *ReceivableDetailServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.ReceivableDetail)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.ReceivableDetailRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service *ReceivableDetailServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	receivableDetailRepo, err := service.ReceivableDetailRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", receivableDetailRepo), err
}

func (service *ReceivableDetailServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	receivableDetailRepo, err := service.ReceivableDetailRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", receivableDetailRepo), err
}

func (service *ReceivableDetailServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params,_ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))

	receivableDetailRepo, totalData, totalFiltered, _ := service.ReceivableDetailRepository.Datatable(ctx, tx, draw, limit, start, search)

	data := make([]interface{}, 0)
	for _, v := range receivableDetailRepo {
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

