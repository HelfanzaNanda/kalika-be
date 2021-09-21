package services

import (
	"strings"

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
	DivisionService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
		Datatable(ctx echo.Context) (res web.Datatable, err error)
	}

	DivisionServiceImpl struct {
		DivisionRepository repository.DivisionRepository
		db *gorm.DB
	}
)

func NewDivisionService(DivisionRepository repository.DivisionRepository, db *gorm.DB) DivisionService {
	return &DivisionServiceImpl{
		DivisionRepository: DivisionRepository,
		db: db,
	}
}

func (service *DivisionServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.Division)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	divisionRepo, err := service.DivisionRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", divisionRepo), err
}

func (service DivisionServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Division)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	divisionRepo, err := service.DivisionRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", divisionRepo), err
}

func (service DivisionServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Division)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.DivisionRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service DivisionServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	divisionRepo, err := service.DivisionRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", divisionRepo), err
}

func (service DivisionServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	divisionRepo, err := service.DivisionRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", divisionRepo), err
}

func (service *DivisionServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params,_ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))

	divisionRepo, totalData, totalFiltered, err := service.DivisionRepository.Datatable(ctx, tx, draw, limit, start, search)
	if err != nil {
		//return helpers.Response(err.Error(), "", nil), err
	}

	var data []interface{}
	for _, v := range divisionRepo {
		v.Action = `<div class="flex">`
		v.Action += `<button type="button" class="flex mr-3" id="edit-data" data-id=`+helpers.IntToString(v.Id)+`> <i data-feather="check-square" class="w-4 h-4 mr-1"></i> Edit </button>`
		v.Action += `<button type="button" class="flex text-theme-6" id="delete-data" data-id=`+helpers.IntToString(v.Id)+`> <i data-feather="trash-2" class="w-4 h-4 mr-1"></i> Delete </button>`
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

//columns := map[int]string{
//	0: "divisions.id",
//	1: "divisions.name",
//}
//var dataOrder []map[string]string
//order := map[string]string{}
//if helpers.ParseFormCollection(ctx.Request(), "order") != nil {
//	order = helpers.ParseFormCollection(ctx.Request(), "order")[0]
//	for k, v := range order {
//		subDataOrder := map[string]string{}
//		if string(k) == "column" {
//			subDataOrder["column"] = columns[helpers.StringToInt(v)]
//			dataOrder = append(dataOrder, subDataOrder)
//		}
//		if string(k) == "dir" {
//			subDataOrder["dir"] = v
//			dataOrder = append(dataOrder, subDataOrder)
//		}
//	}
//}