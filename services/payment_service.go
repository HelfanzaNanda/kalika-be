package services

import (
	//"fmt"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"strings"

	//"kalika-be/config"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"
	"kalika-be/repository"
)

type (
	PaymentService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
		Datatable(ctx echo.Context) (res web.Datatable, err error)
		GeneratePdf(ctx echo.Context) (web.Response, error)
	}

	PaymentServiceImpl struct {
		PaymentRepository repository.PaymentRepository
		db *gorm.DB
	}
)

func NewPaymentService(PaymentRepository repository.PaymentRepository, db *gorm.DB) PaymentService {
	return &PaymentServiceImpl{
		PaymentRepository: PaymentRepository,
		db: db,
	}
}

func (service *PaymentServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.Payment)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	paymentRepo, err := service.PaymentRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", paymentRepo), err
}

func (service *PaymentServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Payment)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	paymentRepo, err := service.PaymentRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", paymentRepo), err
}

func (service *PaymentServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Payment)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.PaymentRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service *PaymentServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	paymentRepo, err := service.PaymentRepository.FindById(ctx, tx, "id", helpers.IntToString(id), map[string][]string{})

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", paymentRepo), err
}

func (service *PaymentServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	paymentRepo, err := service.PaymentRepository.FindAll(ctx, tx, map[string][]string{})

	return helpers.Response("OK", "Sukses Mengambil Data", paymentRepo), err
}

func (service *PaymentServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params,_ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))
	filter := make(map[string]string)
	filter["start_date"] = strings.TrimSpace(params.Get("start_date"))
	filter["end_date"] = strings.TrimSpace(params.Get("end_date"))
	filter["store_id"] = strings.TrimSpace(params.Get("store_id"))
	filter["created_by"] = strings.TrimSpace(params.Get("created_by"))

	paymentRepo, totalData, totalFiltered, _ := service.PaymentRepository.Datatable(ctx, tx, draw, limit, start, search, filter)

	data := make([]interface{}, 0)
	for _, v := range paymentRepo {
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


func (service PaymentServiceImpl) GeneratePdf(ctx echo.Context) (res web.Response, err error) {
	o := new(web.PaymentReportFilterDatatable)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)
	
	paymentRepo, err := service.PaymentRepository.FindByCreatedAt(ctx, tx, o)
	var datas [][]string
	var total float64 = 0
	for _, item := range paymentRepo {
		froot := []string{}
		froot = append(froot, item.Number)
		if item.Model == "PurchaseOrder" {
			froot = append(froot, "Order Pembelian")	
		}else if item.Model == "SalesConsignment" {
			froot = append(froot, "Penjualan Konsinyasi")
		}else if item.Model == "Sales" {
			froot = append(froot, "Penjualan")
		}else if item.Model == "CustomOrder" {
			froot = append(froot, "Penjualan Pesanan")
		}else {
			froot = append(froot, item.Model)
		}
		froot = append(froot, helpers.FormatRupiah(item.Total))
		froot = append(froot, item.PaymentMethod)
		froot = append(froot, item.Date.Format("02 Jan 2006 15:04:05"))
		froot = append(froot, item.CreatedByName)
		datas = append(datas, froot)
		total += item.Total
	}
	title := "laporan_pembayaran"
	headings := []string{"No. Ref", "Tipe", "Total", "Metode", "Dibuat Pada", "Dibuat Oleh"}
	footer := map[string]float64{}
	footer["Total"] = total
	resultPdf, err := helpers.GeneratePdf(ctx, title, headings, datas, footer, o.StartDate, o.EndDate)
	
	return helpers.Response("OK", "Sukses Export PDF", resultPdf), err
}