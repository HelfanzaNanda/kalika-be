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
	SalesReturnService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
		Datatable(ctx echo.Context) (res web.Datatable, err error)
		ReportDatatable(ctx echo.Context) (res web.Datatable, err error)
		GeneratePdf(ctx echo.Context) (web.Response, error)
	}

	SalesReturnServiceImpl struct {
		SalesReturnRepository repository.SalesReturnRepository
		SalesReturnDetailRepository repository.SalesReturnDetailRepository
		db *gorm.DB
	}
)

func NewSalesReturnService(SalesReturnRepository repository.SalesReturnRepository, SalesReturnDetailRepository repository.SalesReturnDetailRepository, db *gorm.DB) SalesReturnService {
	return &SalesReturnServiceImpl{
		SalesReturnRepository: SalesReturnRepository,
		SalesReturnDetailRepository: SalesReturnDetailRepository,
		db: db,
	}
}

func (service *SalesReturnServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(web.SalesReturnPost)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)
	salesReturnRepo, err := service.SalesReturnRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "create sales return error", nil), err
	}
	o.SalesReturn = salesReturnRepo
	salesReturnDetailRepo, err := service.SalesReturnDetailRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "create sales return detail error", nil), err
	}
	o.Total = salesReturnDetailRepo.Total
	
	_, err = service.SalesReturnRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "update total expense error", nil), err
	}
	o.SalesReturnDetails = salesReturnDetailRepo.SalesReturnDetails
	

	return helpers.Response("CREATED", "Sukses Menyimpan Data", o), err
}

func (service SalesReturnServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(web.SalesReturnPost)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.SalesReturnDetailRepository.DeleteBySalesReturnId(ctx, tx, id)
	if err != nil {
		return helpers.Response(err.Error(), "delete sales return detail by sales_return_id error", nil), err
	}
	o.Id = id
	salesReturnDetailRepo, err := service.SalesReturnDetailRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "create sales return detail error", nil), err
	}
	o.Total = salesReturnDetailRepo.Total
	salesReturnRepo, err := service.SalesReturnRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "update sales return error", nil), err
	}
	o.SalesReturn = salesReturnRepo
	o.SalesReturnDetails = salesReturnDetailRepo.SalesReturnDetails

	return helpers.Response("OK", "Sukses Mengubah Data", o), err
}

func (service SalesReturnServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.SalesReturn)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.SalesReturnDetailRepository.DeleteBySalesReturnId(ctx, tx, id)
	if err != nil {
		return helpers.Response(err.Error(), "delete sales return detail error", nil), err
	}
	_, err = service.SalesReturnRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "delete sales return error", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service SalesReturnServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	salesReturn := web.SalesReturnPost{}
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	salesReturnRepo, err := service.SalesReturnRepository.FindById(ctx, tx, "id", helpers.IntToString(id))
	if err != nil {
		return helpers.Response(err.Error(), "find sales return error", nil), err
	}
	salesReturnDetailRepo, err := service.SalesReturnDetailRepository.FindBySalesReturnId(ctx, tx, id)
	if err != nil {
		return helpers.Response(err.Error(), "find sales return detail error", nil), err
	}

	salesReturn.SalesReturn = salesReturnRepo
	salesReturn.SalesReturnDetails = salesReturnDetailRepo

	return helpers.Response("OK", "Sukses Mengambil Data", salesReturn), err
}

func (service SalesReturnServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	salesReturnRepo, err := service.SalesReturnRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", salesReturnRepo), err
}

func (service *SalesReturnServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params,_ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))
	salesRepo, totalData, totalFiltered, _ := service.SalesReturnRepository.Datatable(ctx, tx, draw, limit, start, search)
	// if err != nil {
	// 	return helpers.Response(err.Error(), "", nil), err
	// }

	data := make([]interface{}, 0)
	for _, v := range salesRepo {
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

func (service *SalesReturnServiceImpl) ReportDatatable(ctx echo.Context) (res web.Datatable, err error) {
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
	salesRepo, totalData, totalFiltered, _ := service.SalesReturnRepository.ReportDatatable(ctx, tx, draw, limit, start, search, filter)
	// if err != nil {
	// 	return helpers.Response(err.Error(), "", nil), err
	// }

	data := make([]interface{}, 0)
	for _, v := range salesRepo {
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

func (service SalesReturnServiceImpl) GeneratePdf(ctx echo.Context) (res web.Response, err error) {
	o := new(web.DateRange)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)
	
	salesReturnRepo, err := service.SalesReturnRepository.FindByCreatedAt(ctx, tx, o)
	var datas [][]string
	var total float64 = 0
	for _, item := range salesReturnRepo {
		froot := []string{}
		froot = append(froot, item.Number)
		froot = append(froot, item.CustomerName)
		froot = append(froot, item.StoreConsignmentName)
		froot = append(froot, helpers.FormatRupiah(item.Total))
		froot = append(froot, item.CreatedByName)
		datas = append(datas, froot)

		total += item.Total
	}
	title := "laporan-retur-penjualan"
	headings := []string{"No. Ref", "Kustomer", "Konsiyasi", "Total", "Dibuat Oleh"}
	footer := map[string]float64{}
	footer["Total"] = total
	resultPdf, err := helpers.GeneratePdf(ctx, title, headings, datas, footer)
	
	return helpers.Response("OK", "Sukses Export PDF", resultPdf), err
}