package services

import (
	//"fmt"
	"strings"
	"time"

	// "time"

	"github.com/labstack/echo"
	"gorm.io/gorm"

	//"kalika-be/config"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"
	"kalika-be/repository"
)

type (
	CustomOrderService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
		Datatable(ctx echo.Context) (res web.Datatable, err error)
		ReportDatatable(ctx echo.Context) (res web.Datatable, err error)
		GeneratePdf(ctx echo.Context) (web.Response, error)
	}

	CustomOrderServiceImpl struct {
		CustomOrderRepository repository.CustomOrderRepository
		ReceivableRepository repository.ReceivableRepository
		db                    *gorm.DB
	}
)

func NewCustomOrderService(CustomOrderRepository repository.CustomOrderRepository, ReceivableRepository repository.ReceivableRepository, db *gorm.DB) CustomOrderService {
	return &CustomOrderServiceImpl{
		CustomOrderRepository: CustomOrderRepository,
		ReceivableRepository: ReceivableRepository,
		db:                    db,
	}
}

func (service *CustomOrderServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	receivableRepo := web.ReceivablePosPost{}
	o := new(domain.CustomOrder)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	o.Number = "CO"+helpers.IntToString(int(time.Now().Unix()))
	o.Status = "submission"
	o.CreatedBy = helpers.StringToInt(ctx.Get("userInfo").(map[string]interface{})["id"].(string))

	customOrderRepo, err := service.CustomOrderRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	receivableRepo.Model = "CustomOrder"
	receivableRepo.ModelId = customOrderRepo.Id
	receivableRepo.Total = o.Price
	receivableRepo.Receivables = o.Price
	receivableRepo.Date = time.Now().Format("2006-01-02")
	receivableRepo.Note = o.Number

	service.ReceivableRepository.Create(ctx, tx, &receivableRepo)

	return helpers.Response("CREATED", "Sukses Menyimpan Data", customOrderRepo), err
}

func (service CustomOrderServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.CustomOrder)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	customOrderRepo, err := service.CustomOrderRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", customOrderRepo), err
}

func (service CustomOrderServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.CustomOrder)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.CustomOrderRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service CustomOrderServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	customOrderRepo, err := service.CustomOrderRepository.FindById(ctx, tx, "id", helpers.IntToString(id))
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	completeCustomOrderRepo, err := service.CustomOrderRepository.CompletingResponse(ctx, tx, &customOrderRepo)

	return helpers.Response("OK", "Sukses Mengambil Data", completeCustomOrderRepo), err
}

func (service CustomOrderServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	customOrderRepo, err := service.CustomOrderRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", customOrderRepo), err
}

func (service *CustomOrderServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params, _ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))

	customOrderRepo, totalData, totalFiltered, _ := service.CustomOrderRepository.Datatable(ctx, tx, draw, limit, start, search)

	data := make([]interface{}, 0)
	for _, v := range customOrderRepo {
		v.Action = `<div class="flex">`
		v.Action += `<button type="button" class="btn-edit flex mr-3" id="edit-data" data-id=` + helpers.IntToString(v.Id) + `> <i data-feather="check-square" class="w-4 h-4 mr-1"></i> Edit </button>`
		v.Action += `<button type="button" class="btn-edit flex mr-4 text-theme-12" id="print-data" data-id=`+helpers.IntToString(v.Id)+`> <i data-feather="printer" class="w-4 h-4 mr-1"></i> Cetak</button>`
		v.Action += `<button type="button" class="btn-delete flex text-theme-6" id="delete-data" data-id=` + helpers.IntToString(v.Id) + `> <i data-feather="trash-2" class="w-4 h-4 mr-1"></i> Delete </button>`
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
func (service *CustomOrderServiceImpl) ReportDatatable(ctx echo.Context) (res web.Datatable, err error) {
	params, _ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))
	filter := make(map[string]string)
	filter["start_date"] = strings.TrimSpace(params.Get("start_date"))
	filter["end_date"] = strings.TrimSpace(params.Get("end_date"))
	filter["created_by"] = strings.TrimSpace(params.Get("created_by"))
	filter["payment_method_id"] = strings.TrimSpace(params.Get("payment_method_id"))
	customOrderRepo, totalData, totalFiltered, _ := service.CustomOrderRepository.ReportDatatable(ctx, tx, draw, limit, start, search, filter)

	data := make([]interface{}, 0)
	for _, v := range customOrderRepo {
		data = append(data, v)
	}
	res.Data = data
	res.Order = helpers.ParseFormCollection(ctx.Request(), "order")
	res.Draw = helpers.StringToInt(draw)
	res.RecordsFiltered = totalFiltered
	res.RecordsTotal = totalData

	return res, nil
}

func (service CustomOrderServiceImpl) GeneratePdf(ctx echo.Context) (res web.Response, err error) {
	o := new(web.CustomOrderReportFilterDatatable)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	salesRepo, err := service.CustomOrderRepository.FindByCreatedAt(ctx, tx, o)
	var datas [][]string
	filter := make(map[string]string)
	filter["payment_method_id"] = ""
	if o.PaymentMethodId != 0 {
		filter["payment_method_id"] = helpers.IntToString(o.PaymentMethodId)
	}
	var total float64 = 0
	for _, item := range salesRepo {
		froot := []string{}
			froot = append(froot, item.Number)
			froot = append(froot, item.CreatedAt.Format("02 Jan 2006 15:04:05"))
			froot = append(froot, item.CreatedByName)
			froot = append(froot, helpers.FormatRupiah(item.Total))
			froot = append(froot, item.PaymentMethodName)
			datas = append(datas, froot)
			total += item.Total
		
	}
	title := "laporan-penjualan-pesanan"
	headings := []string{"No. Ref", "Dibuat Pada", "Dibuat Oleh", "Total", "Metode"}
	footer := map[string]float64{}
	footer["Total"] = total
	resultPdf, err := helpers.GeneratePdf(ctx, title, headings, datas, footer)

	return helpers.Response("OK", "Sukses Export PDF", resultPdf), err
}
