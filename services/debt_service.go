package services

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"
	"kalika-be/repository"
	"strings"
)

type (
	DebtService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
		Datatable(ctx echo.Context) (res web.Datatable, err error)
		ReportDatatable(ctx echo.Context) (res web.Datatable, err error)
		GeneratePdf(ctx echo.Context) (web.Response, error)
	}

	DebtServiceImpl struct {
		DebtRepository repository.DebtRepository
		db *gorm.DB
	}
)

func NewDebtService(DebtRepository repository.DebtRepository, db *gorm.DB) DebtService {
	return &DebtServiceImpl{
		DebtRepository: DebtRepository,
		db: db,
	}
}

func (service *DebtServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	debt := domain.Debt{}
	o := new(web.DebtPosPost)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if o.Id > 0 {
		debt, err = service.DebtRepository.Update(ctx, tx, o)
	} else {
		debt, err = service.DebtRepository.Create(ctx, tx, o)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", debt), err
}

func (service DebtServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(web.DebtPosPost)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	debtRepo, err := service.DebtRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", debtRepo), err
}

func (service DebtServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Debt)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.DebtRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service DebtServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	debtRepo, err := service.DebtRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", debtRepo), err
}

func (service DebtServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	debtRepo, err := service.DebtRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", debtRepo), err
}

func (service *DebtServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params,_ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))

	debtRepo, totalData, totalFiltered, _ := service.DebtRepository.Datatable(ctx, tx, draw, limit, start, search)
	// if err != nil {
	// 	return helpers.Response(err.Error(), "", nil), err
	// }

	data := make([]interface{}, 0)
	for _, v := range debtRepo {
		v.Action = `<div class="flex">`
		v.Action += `<button type="button" class="button px-2 mr-1 mb-2 bg-theme-1 text-white" id="pay-data" data-id=`+helpers.IntToString(v.Id)+`> <i data-feather="dollar-sign" class="w-4 h-4 mr-1"></i></button>`
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

func (service *DebtServiceImpl) ReportDatatable(ctx echo.Context) (res web.Datatable, err error) {
	params,_ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))
	filter :=  make(map[string]string)
	filter["start_date"] = strings.TrimSpace(params.Get("start_date"))
	filter["end_date"] = strings.TrimSpace(params.Get("end_date"))
	debtRepo, totalData, totalFiltered, _ := service.DebtRepository.ReportDatatable(ctx, tx, draw, limit, start, search, filter)
	// if err != nil {
	// 	return helpers.Response(err.Error(), "", nil), err
	// }

	data := make([]interface{}, 0)
	for _, v := range debtRepo {
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


func (service DebtServiceImpl) GeneratePdf(ctx echo.Context) (res web.Response, err error) {
	o := new(web.DateRange)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)
	
	debtRepo, err := service.DebtRepository.FindByCreatedAt(ctx, tx, o)
	var datas [][]string
	var totalDebt float64 = 0
	var remainingDebt float64 = 0
	for _, item := range debtRepo {
		froot := []string{}
		froot = append(froot, item.SupplierName)
		froot = append(froot, helpers.FormatRupiah(item.Total))
		froot = append(froot, helpers.FormatRupiah(item.Debts))
		froot = append(froot, item.Date.Format("02 Jan 2006 15:04:05"))
		froot = append(froot, item.Note)
		froot = append(froot, item.CreatedByName)
		datas = append(datas, froot)

		totalDebt += item.Total
		remainingDebt += item.Debts
	}
	title := "laporan_hutang"
	headings := []string{"Supplier", "Total Hutang", "Sisa Hutang", "Tanggal", "Note", "Dibuat Oleh"}
	footer := map[string]float64{}
	footer["Total Hutang"] = totalDebt
	footer["Sisa Hutang"] = remainingDebt
	resultPdf, err := helpers.GeneratePdf(ctx, title, headings, datas, footer, o.StartDate, o.EndDate)
	
	return helpers.Response("OK", "Sukses Export PDF", resultPdf), err
}