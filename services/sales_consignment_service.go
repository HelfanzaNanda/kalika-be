package services

import (
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
	SalesConsignmentService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
		Datatable(ctx echo.Context) (res web.Datatable, err error)
	}

	SalesConsignmentServiceImpl struct {
		SalesConsignmentRepository repository.SalesConsignmentRepository
		SalesConsignmentDetailRepository repository.SalesConsignmentDetailRepository
		PaymentRepository repository.PaymentRepository
		StoreConsignmentRepository repository.StoreConsignmentRepository
		db *gorm.DB
	}
)

func NewSalesConsignmentService(
	SalesConsignmentRepository repository.SalesConsignmentRepository,
	SalesConsignmentDetailRepository repository.SalesConsignmentDetailRepository,
	PaymentRepository repository.PaymentRepository,
	StoreConsignmentRepository repository.StoreConsignmentRepository,
	db *gorm.DB) SalesConsignmentService {
	return &SalesConsignmentServiceImpl{
		SalesConsignmentRepository: SalesConsignmentRepository,
		PaymentRepository: PaymentRepository,
		SalesConsignmentDetailRepository: SalesConsignmentDetailRepository,
		StoreConsignmentRepository: StoreConsignmentRepository,
		db: db,
	}
}

func (service *SalesConsignmentServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	message := "Sukses Menyimpan Data"
	o := new(web.SalesConsignmentPost)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	salesConsignmentRepo := domain.SalesConsignment{}
	salesConsignmentDetailRepo := []web.SalesConsignmentDetailGet{}
	paymentRepo := domain.Payment{}
	storeConsignmentRepo := domain.StoreConsignment{}

	o.SalesConsignment.Total = 0
	for _, val := range o.SalesConsignmentDetails {
		o.SalesConsignment.Total += val.Qty * val.UnitPrice
	}
	o.SalesConsignment.CreatedBy = helpers.StringToInt(ctx.Get("userInfo").(map[string]interface{})["id"].(string))
	if o.Id > 0 {
		message = "Sukses Memperbarui Data"
		salesConsignmentRepo, err = service.SalesConsignmentRepository.Update(ctx, tx, &o.SalesConsignment)
	} else {
		o.SalesConsignment.Number = "SC"+helpers.IntToString(int(time.Now().Unix()))
		o.SalesConsignment.Date = time.Now()
		o.SalesConsignment.Status = "submission"
		salesConsignmentRepo, err = service.SalesConsignmentRepository.Create(ctx, tx, &o.SalesConsignment)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}
	o.SalesConsignment = salesConsignmentRepo

	if o.Id > 0 {
		service.SalesConsignmentDetailRepository.DeleteBySalesConsignment(ctx, tx, o.Id)
		salesConsignmentDetailRepo, err = service.SalesConsignmentDetailRepository.Create(ctx, tx, o)
	} else {
		salesConsignmentDetailRepo, err = service.SalesConsignmentDetailRepository.Create(ctx, tx, o)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}
	o.SalesConsignmentDetails = salesConsignmentDetailRepo

	storeConsignmentRepo, err = service.StoreConsignmentRepository.FindById(ctx, tx, "id", helpers.IntToString(o.StoreConsignmentId))
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	if o.Payment.Id > 0 {
		paymentRepo, err = service.PaymentRepository.Update(ctx, tx, &o.Payment)
	} else {
		o.Payment.Model = "SalesConsignment"
		o.Payment.PaymentMethodId = 12
		o.Payment.Total = o.Total
		o.Payment.PaymentNote = "Piutang Konsinyasi Toko "+storeConsignmentRepo.StoreName
		o.Payment.ModelId = o.Id
		o.Payment.StoreId = helpers.StringToInt(ctx.Get("userInfo").(map[string]interface{})["store_id"].(string))
		o.Payment.CreatedBy = helpers.StringToInt(ctx.Get("userInfo").(map[string]interface{})["id"].(string))
		o.Payment.Date = time.Now()
		paymentRepo, err = service.PaymentRepository.Create(ctx, tx, &o.Payment)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}
	o.Payment = paymentRepo

	return helpers.Response("CREATED", message, o), err
}

func (service *SalesConsignmentServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.SalesConsignment)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	salesConsignmentRepo, err := service.SalesConsignmentRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", salesConsignmentRepo), err
}

func (service *SalesConsignmentServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.SalesConsignment)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.SalesConsignmentRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service *SalesConsignmentServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	salesConsignment := web.SalesConsignmentPost{}
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	salesConsignmentRepo, err := service.SalesConsignmentRepository.FindById(ctx, tx, "id", helpers.IntToString(id))
	paymentSearch := make(map[string][]string)
	paymentSearch["model_id"] = append(paymentSearch["model_id"], helpers.IntToString(salesConsignmentRepo.Id))
	paymentRepo, err := service.PaymentRepository.FindById(ctx, tx, "model", "SalesConsignment", paymentSearch)
	salesConsignmentDetailSearch := make(map[string][]string)
	salesConsignmentDetailSearch["sales_consignment_id"] = append(paymentSearch["model_id"], helpers.IntToString(salesConsignmentRepo.Id))
	salesConsignmentDetailRepo, err := service.SalesConsignmentDetailRepository.FindAll(ctx, tx, salesConsignmentDetailSearch)

	salesConsignment.SalesConsignment = salesConsignmentRepo
	salesConsignment.Payment = paymentRepo
	salesConsignment.SalesConsignmentDetails = salesConsignmentDetailRepo

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", salesConsignment), err
}

func (service *SalesConsignmentServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	salesConsignmentRepo, err := service.SalesConsignmentRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", salesConsignmentRepo), err
}

func (service *SalesConsignmentServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params,_ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))

	recipeRepo, totalData, totalFiltered, _ := service.SalesConsignmentRepository.Datatable(ctx, tx, draw, limit, start, search)

	data := make([]interface{}, 0)
	for _, v := range recipeRepo {
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

