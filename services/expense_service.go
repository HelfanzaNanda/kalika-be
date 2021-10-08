package services

import (
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
	ExpenseService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
		Datatable(ctx echo.Context) (res web.Datatable, err error)
		ReportDatatable(ctx echo.Context) (res web.Datatable, err error)
	}

	ExpenseServiceImpl struct {
		ExpenseRepository repository.ExpenseRepository
		ExpenseDetailRepository repository.ExpenseDetailRepository
		db *gorm.DB
	}
)

func NewExpenseService(ExpenseRepository repository.ExpenseRepository, ExpenseDetailRepository repository.ExpenseDetailRepository, db *gorm.DB) ExpenseService {
	return &ExpenseServiceImpl{
		ExpenseRepository: ExpenseRepository,
		ExpenseDetailRepository: ExpenseDetailRepository,
		db: db,
	}
}

func (service *ExpenseServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(web.ExpensePosPost)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)
	expenseRepo, err := service.ExpenseRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "create expense error", nil), err
	}
	o.Expense = expenseRepo

	expenseDetailRepo, err := service.ExpenseDetailRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "create expense detail error", nil), err
	}
	o.Total = expenseDetailRepo.Total
	
	_, err = service.ExpenseRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "update total expense error", nil), err
	}
	o.ExpenseDetails = expenseDetailRepo.ExpenseDetails
	

	return helpers.Response("CREATED", "Sukses Menyimpan Data", o), err
}

func (service ExpenseServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(web.ExpensePosPost)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.ExpenseDetailRepository.DeleteByExpenseId(ctx, tx, id)
	if err != nil {
		return helpers.Response(err.Error(), "delete expense detail by expense_id error", nil), err
	}
	o.Id = id
	expenseDetailRepo, err := service.ExpenseDetailRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "create many expense detail error", nil), err
	}
	o.Total = expenseDetailRepo.Total
	expenseRepo, err := service.ExpenseRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "update expense error", nil), err
	}
	o.Expense = expenseRepo
	o.ExpenseDetails = expenseDetailRepo.ExpenseDetails

	return helpers.Response("OK", "Sukses Mengubah Data", o), err
}

func (service ExpenseServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Expense)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)
	_, err = service.ExpenseDetailRepository.DeleteByExpenseId(ctx, tx, id)
	if err != nil {
		return helpers.Response(err.Error(), "delete expense detail error", nil), err
	}
	_, err = service.ExpenseRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}


	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service ExpenseServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	expensePost := web.ExpensePosPost{}
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	expenseRepo, err := service.ExpenseRepository.FindById(ctx, tx, "id", helpers.IntToString(id))
	
	expenseDetailRepo, err := service.ExpenseDetailRepository.FindByExpenseId(ctx, tx, id)

	expensePost.Expense = expenseRepo
	expensePost.ExpenseDetails = expenseDetailRepo

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", expensePost), err
}

func (service ExpenseServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	expenseRepo, err := service.ExpenseRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", expenseRepo), err
}

func (service *ExpenseServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params,_ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))

	expenseRepo, totalData, totalFiltered, _ := service.ExpenseRepository.Datatable(ctx, tx, draw, limit, start, search)
	// if err != nil {
	// 	return helpers.Response(err.Error(), "", nil), err
	// }

	data := make([]interface{}, 0)
	for _, v := range expenseRepo {
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

func (service *ExpenseServiceImpl) ReportDatatable(ctx echo.Context) (res web.Datatable, err error) {
	params,_ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))

	expenseRepo, totalData, totalFiltered, _ := service.ExpenseRepository.ReportDatatable(ctx, tx, draw, limit, start, search)
	// if err != nil {
	// 	return helpers.Response(err.Error(), "", nil), err
	// }

	data := make([]interface{}, 0)
	for _, v := range expenseRepo {
		v.Action = `<div class="flex">`
		v.Action += `<button type="button" class="btn-edit flex mr-3" id="edit-data" data-id=`+helpers.IntToString(v.Id)+`> <i data-feather="pdf" class="w-4 h-4 mr-1"></i> Print </button>`
		// v.Action += `<button type="button" class="btn-delete flex text-theme-6" id="delete-data" data-id=`+helpers.IntToString(v.Id)+`> <i data-feather="trash-2" class="w-4 h-4 mr-1"></i> Delete </button>`
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