package services

import (
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
	SalesService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
		Datatable(ctx echo.Context) (res web.Datatable, err error)
		ReportDatatable(ctx echo.Context) (res web.Datatable, err error)
		GeneratePdf(ctx echo.Context) (web.Response, error)
	}

	SalesServiceImpl struct {
		SalesRepository       repository.SalesRepository
		SalesDetailRepository repository.SalesDetailRepository
		PaymentRepository repository.PaymentRepository
		CustomerRepository repository.CustomerRepository
		ProductLocationRepository repository.ProductLocationRepository
		ReceivableRepository repository.ReceivableRepository
		db *gorm.DB
	}
)

func NewSalesService(SalesRepository repository.SalesRepository, SalesDetailRepository repository.SalesDetailRepository, PaymentRepository repository.PaymentRepository, CustomerRepository repository.CustomerRepository, ProductLocationRepository repository.ProductLocationRepository, ReceivableRepository repository.ReceivableRepository, db *gorm.DB) SalesService {
	return &SalesServiceImpl{
		SalesRepository:       SalesRepository,
		SalesDetailRepository: SalesDetailRepository,
		PaymentRepository: PaymentRepository,
		CustomerRepository: CustomerRepository,
		ProductLocationRepository: ProductLocationRepository,
		ReceivableRepository: ReceivableRepository,
		db: db,
	}
}

func (service *SalesServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	message := "Sukses Menyimpan Data"
	o := new(web.SalesPosPost)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	customerRepo := domain.Customer{}
	salesRepo := domain.Sale{}
	salesDetailRepo := []web.SalesDetailPosGet{}
	paymentRepo := domain.Payment{}
	receivableRepo := web.ReceivablePosPost{}

	if o.Customer.Id > 0 {
		customerRepo, err = service.CustomerRepository.Update(ctx, tx, &o.Customer)
	} else if o.Customer.Name != "" || o.Customer.Phone != "" {
		customerRepo, err = service.CustomerRepository.Create(ctx, tx, &o.Customer)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}
	o.Customer = customerRepo

	if o.Id > 0 {
		message = "Sukses Memperbarui Data"
		salesRepo, err = service.SalesRepository.Update(ctx, tx, &o.Sale)
	} else {
		salesRepo, err = service.SalesRepository.Create(ctx, tx, o)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}
	o.Sale = salesRepo

	if o.Id > 0 {
		service.SalesDetailRepository.DeleteBySales(ctx, tx, o.Id)
		salesDetailRepo, err = service.SalesDetailRepository.Create(ctx, tx, o)
	} else {
		salesDetailRepo, err = service.SalesDetailRepository.Create(ctx, tx, o)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}
	o.SalesDetails = salesDetailRepo

	if o.Payment.Id > 0 {
		paymentRepo, err = service.PaymentRepository.Update(ctx, tx, &o.Payment)
	} else {
		o.Payment.Model = "Sales"
		o.Payment.ModelId = o.Id
		o.Payment.StoreId = helpers.StringToInt(ctx.Get("userInfo").(map[string]interface{})["store_id"].(string))
		o.Payment.CreatedBy = helpers.StringToInt(ctx.Get("userInfo").(map[string]interface{})["id"].(string))
		o.Payment.Date = time.Now()
		paymentRepo, err = service.PaymentRepository.Create(ctx, tx, &o.Payment)

		if o.Payment.Total > o.CustomerPay {
			// Receivable
			receivableRepo.Model = "Sales"
			receivableRepo.ModelId = salesRepo.Id
			receivableRepo.Total = o.Payment.Total - o.CustomerPay
			receivableRepo.Receivables = o.Payment.Total - o.CustomerPay
			receivableRepo.Date = time.Now().Format("2006-01-02")
			receivableRepo.Note = o.Number

			service.ReceivableRepository.Create(ctx, tx, &receivableRepo)
		}

		productLocations := []map[string]interface{}{}
		for _, val := range salesDetailRepo {
			productLocation := map[string]interface{}{}
			productLocation["model"] = "Product"
			productLocation["product_id"] = helpers.IntToString(val.ProductId)
			productLocation["quantity"] = helpers.IntToString(val.Qty)
			productLocation["store_id"] = helpers.IntToString(salesRepo.StoreId)
			productLocations = append(productLocations, productLocation)
		}

		service.ProductLocationRepository.StockDeduction(ctx, tx, productLocations)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}
	o.Payment = paymentRepo

	return helpers.Response("CREATED", message, o), err
}

func (service SalesServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Sale)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	salesRepo, err := service.SalesRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", salesRepo), err
}

func (service SalesServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Sale)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.SalesRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service SalesServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	salesPos := web.SalesPosPost{}
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	salesRepo, err := service.SalesRepository.FindById(ctx, tx, "id", helpers.IntToString(id))
	customerRepo, err := service.CustomerRepository.FindById(ctx, tx, "id", helpers.IntToString(salesRepo.CustomerId))
	paymentSearch := make(map[string][]string)
	paymentSearch["model_id"] = append(paymentSearch["model_id"], helpers.IntToString(salesRepo.Id))
	paymentRepo, err := service.PaymentRepository.FindById(ctx, tx, "model", "Sales", paymentSearch)
	salesDetailSearch := make(map[string][]string)
	salesDetailSearch["sales_id"] = append(paymentSearch["model_id"], helpers.IntToString(salesRepo.Id))
	salesDetailRepo, err := service.SalesDetailRepository.FindAll(ctx, tx, salesDetailSearch)

	salesPos.Sale = salesRepo
	salesPos.Payment = paymentRepo
	salesPos.Customer = customerRepo
	salesPos.SalesDetails = salesDetailRepo

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", salesPos), err
}

func (service SalesServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	salesRepo, err := service.SalesRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", salesRepo), err
}

func (service *SalesServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params, _ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))

	salesRepo, totalData, totalFiltered, _ := service.SalesRepository.Datatable(ctx, tx, draw, limit, start, search)
	// if err != nil {
	// 	return helpers.Response(err.Error(), "", nil), err
	// }

	data := make([]interface{}, 0)
	for _, v := range salesRepo {
		v.Action = `<div class="flex">`
		v.Action += `<button type="button" class="btn-edit flex mr-3" id="edit-data" data-id=` + helpers.IntToString(v.Id) + `> <i data-feather="check-square" class="w-4 h-4 mr-1"></i> Edit </button>`
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

func (service *SalesServiceImpl) ReportDatatable(ctx echo.Context) (res web.Datatable, err error) {
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
	filter["store_id"] = strings.TrimSpace(params.Get("store_id"))
	filter["created_by"] = strings.TrimSpace(params.Get("created_by"))
	filter["payment_method_id"] = strings.TrimSpace(params.Get("payment_method_id"))
	saleRepo, totalData, totalFiltered, _ := service.SalesRepository.ReportDatatable(ctx, tx, draw, limit, start, search, filter)

	data := make([]interface{}, 0)
	for _, v := range saleRepo {
		paymentMethodName, _ := service.PaymentRepository.FindByModel(ctx, tx, "Sales", v.Id, filter)
		if paymentMethodName.PaymentMethodName != "" {
			v.Metode = paymentMethodName.PaymentMethodName
			data = append(data, v)
		}
	}
	res.Data = data
	res.Order = helpers.ParseFormCollection(ctx.Request(), "order")
	res.Draw = helpers.StringToInt(draw)
	res.RecordsFiltered = totalFiltered
	res.RecordsTotal = totalData

	return res, nil
}

func (service SalesServiceImpl) GeneratePdf(ctx echo.Context) (res web.Response, err error) {
	o := new(web.SaleReportFilterDatatable)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	salesRepo, err := service.SalesRepository.FindByCreatedAt(ctx, tx, o)
	var datas [][]string
	filter := make(map[string]string)
	filter["payment_method_id"] = ""
	if o.PaymentMethodId != 0 {
		filter["payment_method_id"] = helpers.IntToString(o.PaymentMethodId)
	}
	var discount float64 = 0
	var total float64 = 0
	for _, item := range salesRepo {
		paymentMethodName, _ := service.PaymentRepository.FindByModel(ctx, tx, "Sales", item.Id, filter)
		if paymentMethodName.PaymentMethodName != "" {
			froot := []string{}
			//froot = append(froot, helpers.IntToString(index + 1))
			froot = append(froot, item.CreatedAt.Format("02 Jan 2006 15:04:05"))
			froot = append(froot, item.CreatedByName)
			froot = append(froot, item.Number)
			froot = append(froot, helpers.FormatRupiah(float64(item.DiscountValue)))
			froot = append(froot, helpers.FormatRupiah(item.Total))
			froot = append(froot, paymentMethodName.PaymentMethodName)
			datas = append(datas, froot)
			total += item.Total
			total += float64(item.DiscountValue)
		}
		
	}
	title := "laporan-penjualan"
	headings := []string{"Dibuat Pada", "Dibuat Oleh", "No. Ref", "Diskon", "Total", "Metode"}
	footer := map[string]float64{}
	footer["Diskon"] = discount
	footer["Total"] = total
	resultPdf, err := helpers.GeneratePdf(ctx, title, headings, datas, footer)

	return helpers.Response("OK", "Sukses Export PDF", resultPdf), err
}
