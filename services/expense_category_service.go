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
	ExpenseCategoryService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
		Datatable(ctx echo.Context) (res web.Datatable, err error)
	}

	ExpenseCategoryServiceImpl struct {
		ExpenseCategoryRepository repository.ExpenseCategoryRepository
		db *gorm.DB
	}
)

func NewExpenseCategoryService(ExpenseCategoryRepository repository.ExpenseCategoryRepository, db *gorm.DB) ExpenseCategoryService {
	return &ExpenseCategoryServiceImpl{
		ExpenseCategoryRepository: ExpenseCategoryRepository,
		db: db,
	}
}

func (service *ExpenseCategoryServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	expenseCategoryRepo := domain.ExpenseCategory{}
	o := new(domain.ExpenseCategory)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if o.Id > 0 {
		expenseCategoryRepo, err = service.ExpenseCategoryRepository.Update(ctx, tx, o)
	} else {
		expenseCategoryRepo, err = service.ExpenseCategoryRepository.Create(ctx, tx, o)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", expenseCategoryRepo), err
}

func (service ExpenseCategoryServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.ExpenseCategory)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	expenseCategoryRepo, err := service.ExpenseCategoryRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", expenseCategoryRepo), err
}

func (service ExpenseCategoryServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.ExpenseCategory)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.ExpenseCategoryRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service ExpenseCategoryServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	expenseCategoryRepo, err := service.ExpenseCategoryRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", expenseCategoryRepo), err
}

func (service ExpenseCategoryServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	expenseCategoryRepo, err := service.ExpenseCategoryRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", expenseCategoryRepo), err
}

func (service *ExpenseCategoryServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params,_ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))

	expenseCategoryRepo, totalData, totalFiltered, _ := service.ExpenseCategoryRepository.Datatable(ctx, tx, draw, limit, start, search)
	// if err != nil {
	// 	return helpers.Response(err.Error(), "", nil), err
	// }

	data := make([]interface{}, 0)
	for _, v := range expenseCategoryRepo {
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
