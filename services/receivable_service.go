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
	ReceivableService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
		Datatable(ctx echo.Context) (res web.Datatable, err error)
		ReportDatatable(ctx echo.Context) (res web.Datatable, err error)
		GeneratePdf(ctx echo.Context) (web.Response, error)
	}

	ReceivableServiceImpl struct {
		ReceivableRepository repository.ReceivableRepository
		db *gorm.DB
	}
)

func NewReceivableService(ReceivableRepository repository.ReceivableRepository, db *gorm.DB) ReceivableService {
	return &ReceivableServiceImpl{
		ReceivableRepository: ReceivableRepository,
		db: db,
	}
}

func (service *ReceivableServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	receivable := domain.Receivable{}
	o := new(web.ReceivablePosPost)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if o.Id > 0 {
		receivable, err = service.ReceivableRepository.Update(ctx, tx, o)
	} else {
		receivable, err = service.ReceivableRepository.Create(ctx, tx, o)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", receivable), err
}

func (service ReceivableServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(web.ReceivablePosPost)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	receivableRepo, err := service.ReceivableRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", receivableRepo), err
}

func (service ReceivableServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Receivable)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.ReceivableRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service ReceivableServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	receivableRepo, err := service.ReceivableRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", receivableRepo), err
}

func (service ReceivableServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	receivableRepo, err := service.ReceivableRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", receivableRepo), err
}

func (service *ReceivableServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params,_ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))

	receivableRepo, totalData, totalFiltered, _ := service.ReceivableRepository.Datatable(ctx, tx, draw, limit, start, search)

	data := make([]interface{}, 0)
	for _, v := range receivableRepo {
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

func (service *ReceivableServiceImpl) ReportDatatable(ctx echo.Context) (res web.Datatable, err error) {
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
	receivableRepo, totalData, totalFiltered, _ := service.ReceivableRepository.ReportDatatable(ctx, tx, draw, limit, start, search, filter)
	// if err != nil {
	// 	return helpers.Response(err.Error(), "", nil), err
	// }

	data := make([]interface{}, 0)
	for _, v := range receivableRepo {
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

func (service ReceivableServiceImpl) GeneratePdf(ctx echo.Context) (res web.Response, err error) {
	o := new(web.DateRange)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)
	
	receivableRepo, err := service.ReceivableRepository.FindByCreatedAt(ctx, tx, o)
	var datas [][]string
	for _, item := range receivableRepo {
		froot := []string{}
		froot = append(froot, helpers.IntToString(int(item.Total)))
		froot = append(froot, helpers.IntToString(int(item.Receivables)))
		froot = append(froot, item.CustomerName)
		froot = append(froot, item.StoreConsignmentName)
		froot = append(froot, item.Date.String())
		froot = append(froot, item.Note)
		
		datas = append(datas, froot)
	}
	title := "laporan-piutang"
	headings := []string{"Total", "Receivables", "Customer Name", "Store Consignment Name", "Date", "Note"}
	resultPdf, err := helpers.GeneratePdf(ctx, title, headings, datas)
	
	return helpers.Response("OK", "Sukses Export PDF", resultPdf), err
}