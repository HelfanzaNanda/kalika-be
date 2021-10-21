package services

import (
	"fmt"
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
	PurchaseOrderService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
		Datatable(ctx echo.Context) (res web.Datatable, err error)
		ReportDatatable(ctx echo.Context) (res web.Datatable, err error)
		GeneratePdf(ctx echo.Context) (web.Response, error)
	}

	PurchaseOrderServiceImpl struct {
		PurchaseOrderRepository repository.PurchaseOrderRepository
		PurchaseOrderDetailRepository repository.PurchaseOrderDetailRepository
		PaymentRepository repository.PaymentRepository
		ProductLocationRepository repository.ProductLocationRepository
		DebtRepository repository.DebtRepository
		db *gorm.DB
	}
)

func NewPurchaseOrderService(PurchaseOrderRepository repository.PurchaseOrderRepository, PurchaseOrderDetailRepository repository.PurchaseOrderDetailRepository, PaymentRepository repository.PaymentRepository, ProductLocationRepository repository.ProductLocationRepository, DebtRepository repository.DebtRepository, db *gorm.DB) PurchaseOrderService {
	return &PurchaseOrderServiceImpl{
		PurchaseOrderRepository: PurchaseOrderRepository,
		PurchaseOrderDetailRepository: PurchaseOrderDetailRepository,
		PaymentRepository: PaymentRepository,
		ProductLocationRepository: ProductLocationRepository,
		DebtRepository: DebtRepository,
		db: db,
	}
}

func (service *PurchaseOrderServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	message := "Sukses Menyimpan Data"
	o := new(web.PurchaseOrderPost)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	purchaseOrderRepo := domain.PurchaseOrder{}
	purchaseOrderDetailRepo := []web.PurchaseOrderDetailGet{}
	paymentRepo := domain.Payment{}
	debtRepo := web.DebtPosPost{}

	if o.Id > 0 {
		message = "Sukses Memperbarui Data"
		purchaseOrderRepo, err = service.PurchaseOrderRepository.Update(ctx, tx, &o.PurchaseOrder)
	} else {
		o.PurchaseOrder.Number = "PO"+helpers.IntToString(int(time.Now().Unix()))
		o.PurchaseOrder.Date = time.Now()
		o.PurchaseOrder.Status = "submission"
		o.PurchaseOrder.CreatedBy = helpers.StringToInt(ctx.Get("userInfo").(map[string]interface{})["id"].(string))
		purchaseOrderRepo, err = service.PurchaseOrderRepository.Create(ctx, tx, &o.PurchaseOrder)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}
	o.PurchaseOrder = purchaseOrderRepo

	if o.Id > 0 {
		service.PurchaseOrderDetailRepository.DeleteByPurchaseOrder(ctx, tx, o.Id)
		purchaseOrderDetailRepo, err = service.PurchaseOrderDetailRepository.Create(ctx, tx, o)
	} else {
		purchaseOrderDetailRepo, err = service.PurchaseOrderDetailRepository.Create(ctx, tx, o)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}
	o.PurchaseOrderDetails = purchaseOrderDetailRepo

	if o.Payment.Id > 0 {
		paymentRepo, err = service.PaymentRepository.Update(ctx, tx, &o.Payment)
	} else {
		o.Payment.Model = "PurchaseOrder"
		o.Payment.ModelId = o.Id
		o.Payment.StoreId = helpers.StringToInt(ctx.Get("userInfo").(map[string]interface{})["store_id"].(string))
		o.Payment.CreatedBy = helpers.StringToInt(ctx.Get("userInfo").(map[string]interface{})["id"].(string))
		o.Payment.Date = time.Now()
		paymentRepo, err = service.PaymentRepository.Create(ctx, tx, &o.Payment)

		productLocations := []map[string]interface{}{}
		for _, val := range purchaseOrderDetailRepo {
			productLocation := map[string]interface{}{}
			productLocation["model"] = "RawMaterial"
			productLocation["product_id"] = helpers.IntToString(val.RawMaterialId)
			productLocation["quantity"] = helpers.IntToString(val.Qty)
			productLocation["store_id"] = ctx.Get("userInfo").(map[string]interface{})["store_id"].(string)
			productLocations = append(productLocations, productLocation)
		}

		_, err = service.ProductLocationRepository.StockAddition(ctx, tx, productLocations)

		if o.Payment.Total > o.CustomerPay {
			fmt.Println(purchaseOrderRepo)
			debtRepo.Model = "PurchaseOrder"
			debtRepo.ModelId = purchaseOrderRepo.Id
			debtRepo.Total = o.Payment.Total - o.CustomerPay
			debtRepo.Debts = o.Payment.Total - o.CustomerPay
			debtRepo.Date = time.Now().Format("2006-01-02")
			debtRepo.Note = o.PurchaseOrder.Number
			debtRepo.SupplierId = purchaseOrderRepo.SupplierId

			service.DebtRepository.Create(ctx, tx, &debtRepo)
		}
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}
	o.Payment = paymentRepo

	return helpers.Response("CREATED", message, o), err
}

func (service PurchaseOrderServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.PurchaseOrder)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	purchaseOrderRepo, err := service.PurchaseOrderRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", purchaseOrderRepo), err
}

func (service PurchaseOrderServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.PurchaseOrder)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.PurchaseOrderRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service PurchaseOrderServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	purchaseOrder := web.PurchaseOrderPost{}
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	purchaseOrderRepo, err := service.PurchaseOrderRepository.FindById(ctx, tx, "id", helpers.IntToString(id))
	paymentSearch := make(map[string][]string)
	paymentSearch["model_id"] = append(paymentSearch["model_id"], helpers.IntToString(purchaseOrderRepo.Id))
	paymentRepo, err := service.PaymentRepository.FindById(ctx, tx, "model", "PurchaseOrder", paymentSearch)
	purchaseOrderDetailSearch := make(map[string][]string)
	purchaseOrderDetailSearch["purchase_order_id"] = append(paymentSearch["model_id"], helpers.IntToString(purchaseOrderRepo.Id))
	purchaseOrderDetailRepo, err := service.PurchaseOrderDetailRepository.FindAll(ctx, tx, purchaseOrderDetailSearch)

	purchaseOrder.PurchaseOrder = purchaseOrderRepo
	purchaseOrder.Payment = paymentRepo
	purchaseOrder.PurchaseOrderDetails = purchaseOrderDetailRepo

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", purchaseOrder), err
}

func (service PurchaseOrderServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	purchaseOrderRepo, err := service.PurchaseOrderRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", purchaseOrderRepo), err
}

func (service *PurchaseOrderServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params,_ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))
	purchaseOrderRepo, totalData, totalFiltered, _ := service.PurchaseOrderRepository.Datatable(ctx, tx, draw, limit, start, search)

	data := make([]interface{}, 0)
	for _, v := range purchaseOrderRepo {
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


func (service *PurchaseOrderServiceImpl) ReportDatatable(ctx echo.Context) (res web.Datatable, err error) {
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
	purchaseOrderRepo, totalData, totalFiltered, _ := service.PurchaseOrderRepository.ReportDatatable(ctx, tx, draw, limit, start, search, filter)
	// if err != nil {
	// 	return helpers.Response(err.Error(), "", nil), err
	// }

	data := make([]interface{}, 0)
	for _, v := range purchaseOrderRepo {
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

func (service PurchaseOrderServiceImpl) GeneratePdf(ctx echo.Context) (res web.Response, err error) {
	o := new(web.DateRange)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)
	
	purchaseOrderRepo, err := service.PurchaseOrderRepository.FindByCreatedAt(ctx, tx, o)
	var datas [][]string
	var total float64 = 0
	for _, item := range purchaseOrderRepo {
		froot := []string{}
		froot = append(froot, item.Number)
		froot = append(froot, item.SupplierName)
		froot = append(froot, item.Status)
		froot = append(froot, helpers.FormatRupiah(item.Total))
		froot = append(froot, item.CreatedByName)
		froot = append(froot, item.CreatedAt.Format("02 Jan 2006 15:04:05"))
		datas = append(datas, froot)

		total += item.Total
	}
	title := "laporan-pembelian"
	headings := []string{"No. Ref", "Supplier", "Status", "Total", "Dibuat Oleh", "Dibuat Pada"}
	footer := map[string]float64{}
	footer["Total"] = total
	resultPdf, err := helpers.GeneratePdf(ctx, title, headings, datas, footer)
	
	return helpers.Response("OK", "Sukses Export PDF", resultPdf), err
}