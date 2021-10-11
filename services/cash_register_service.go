package services

import (
	//"fmt"
	"strings"
	"time"

	"github.com/labstack/echo"
	"gorm.io/gorm"

	//"kalika-be/config"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"
	"kalika-be/repository"
)

type (
	CashRegisterService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
		Datatable(ctx echo.Context) (res web.Datatable, err error)

	}

	CashRegisterServiceImpl struct {
		CashRegisterRepository repository.CashRegisterRepository
		db *gorm.DB
	}
)

func NewCashRegisterService(CashRegisterRepository repository.CashRegisterRepository, db *gorm.DB) CashRegisterService {
	return &CashRegisterServiceImpl{
		CashRegisterRepository: CashRegisterRepository,
		db: db,
	}
}

func (service *CashRegisterServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	cashRegisterRepo := domain.CashRegister{}
	o := new(domain.CashRegister)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	o.Number = "CR"+helpers.IntToString(int(time.Now().Unix()))
	if o.Id > 0 {
		cashRegisterRepo, err = service.CashRegisterRepository.Update(ctx, tx, o)
	} else {
		cashRegisterRepo, err = service.CashRegisterRepository.Create(ctx, tx, o)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", cashRegisterRepo), err
}

func (service CashRegisterServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.CashRegister)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	cashRegisterRepo, err := service.CashRegisterRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", cashRegisterRepo), err
}

func (service CashRegisterServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.CashRegister)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.CashRegisterRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service CashRegisterServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	cashRegisterRepo, err := service.CashRegisterRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", cashRegisterRepo), err
}

func (service CashRegisterServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	cashRegisterRepo, err := service.CashRegisterRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", cashRegisterRepo), err
}

func (service *CashRegisterServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params,_ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))

	divisionRepo, totalData, totalFiltered, _ := service.CashRegisterRepository.Datatable(ctx, tx, draw, limit, start, search)

	data := make([]interface{}, 0)
	for _, v := range divisionRepo {
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