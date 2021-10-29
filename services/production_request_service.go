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
	ProductionRequestService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
		Datatable(ctx echo.Context) (res web.Datatable, err error)
		GeneratePdf(ctx echo.Context, productionRequestId int) (web.Response, error)
	}

	ProductionRequestServiceImpl struct {
		ProductionRequestRepository repository.ProductionRequestRepository
		ProductionRequestDetailRepository repository.ProductionRequestDetailRepository
		PaymentRepository repository.PaymentRepository
		StoreConsignmentRepository repository.StoreConsignmentRepository
		db *gorm.DB
	}
)

func NewProductionRequestService(
	ProductionRequestRepository repository.ProductionRequestRepository,
	ProductionRequestDetailRepository repository.ProductionRequestDetailRepository,
	db *gorm.DB) ProductionRequestService {
	return &ProductionRequestServiceImpl{
		ProductionRequestRepository: ProductionRequestRepository,
		ProductionRequestDetailRepository: ProductionRequestDetailRepository,
		db: db,
	}
}

func (service *ProductionRequestServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	message := "Sukses Menyimpan Data"
	o := new(web.ProductionRequestPost)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	productionRequestRepo := web.ProductionRequestGet{}
	productionRequestDetailRepo := []web.ProductionRequestDetailGet{}

	o.ProductionRequest.CreatedBy = helpers.StringToInt(ctx.Get("userInfo").(map[string]interface{})["id"].(string))
	if o.Id > 0 {
		message = "Sukses Memperbarui Data"
		productionRequestRepo, err = service.ProductionRequestRepository.Update(ctx, tx, &o.ProductionRequest)
	} else {
		o.ProductionRequest.Number = "PR"+helpers.IntToString(int(time.Now().Unix()))
		o.ProductionRequest.Status = "submission"
		o.ProductionRequest.Date = time.Now()
		productionRequestRepo, err = service.ProductionRequestRepository.Create(ctx, tx, &o.ProductionRequest)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	for key, _ := range o.ProductionRequestDetail {
		o.ProductionRequestDetail[key].ProductionRequestId = productionRequestRepo.Id
	}

	if o.Id > 0 {
		service.ProductionRequestDetailRepository.DeleteByProductionRequest(ctx, tx, o.Id)
		productionRequestDetailRepo, err = service.ProductionRequestDetailRepository.Create(ctx, tx, o.ProductionRequestDetail)
	} else {
		productionRequestDetailRepo, err = service.ProductionRequestDetailRepository.Create(ctx, tx, o.ProductionRequestDetail)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}
	productionRequestRepo.ProductionRequestDetail = productionRequestDetailRepo

	return helpers.Response("CREATED", message, productionRequestRepo), err
}

func (service *ProductionRequestServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.ProductionRequest)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	productionRequestRepo, err := service.ProductionRequestRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", productionRequestRepo), err
}

func (service *ProductionRequestServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.ProductionRequest)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.ProductionRequestRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service *ProductionRequestServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	productionRequest := web.ProductionRequestGet{}
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	productionRequestRepo, err := service.ProductionRequestRepository.FindById(ctx, tx, "id", helpers.IntToString(id))
	productionRequestDetailSearch := make(map[string][]string)
	productionRequestDetailSearch["production_request_id"] = append(productionRequestDetailSearch["production_request_id"], helpers.IntToString(productionRequestRepo.Id))
	productionRequestDetailRepo, err := service.ProductionRequestDetailRepository.FindAll(ctx, tx, productionRequestDetailSearch)

	productionRequest = productionRequestRepo
	productionRequest.ProductionRequestDetail = productionRequestDetailRepo

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", productionRequest), err
}

func (service *ProductionRequestServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	productionRequestRepo, err := service.ProductionRequestRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", productionRequestRepo), err
}

func (service *ProductionRequestServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params,_ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))

	recipeRepo, totalData, totalFiltered, _ := service.ProductionRequestRepository.Datatable(ctx, tx, draw, limit, start, search)

	data := make([]interface{}, 0)
	for _, v := range recipeRepo {
		v.Action = `<div class="flex">`
		v.Action += `<button type="button" class="btn-edit flex " id="edit-data" data-id=`+helpers.IntToString(v.Id)+`> <i data-feather="check-square" class="w-4 h-4 mr-1"></i> Edit </button>`
		v.Action += `<button type="button" class="btn-delete flex mx-3 text-theme-6" id="delete-data" data-id=`+helpers.IntToString(v.Id)+`> <i data-feather="trash-2" class="w-4 h-4 mr-1"></i> Delete </button>`
		v.Action += `<button type="button" class="btn-pdf flex text-theme-6" id="pdf-data" data-id=`+helpers.IntToString(v.Id)+`> <i data-feather="file-text" class="w-4 h-4 mr-1"></i> Pdf </button>`
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

func (service *ProductionRequestServiceImpl) GeneratePdf(ctx echo.Context, productionRequestId int) (res web.Response, err error) {
	//tx := service.db.Begin()
	//defer helpers.CommitOrRollback(tx)
	//
	//productionRequestRepo, err := service.ProductionRequestDetailRepository.Pdf(ctx, tx, productionRequestId)
	//var datas [][]string
	//for _, item := range productionRequestRepo {
	//	froot := []string{}
	//		froot = append(froot, item.ProductName)
	//		froot = append(froot, item.CategoryName)
	//		froot = append(froot, helpers.IntToString(item.MinimumStock))
	//		froot = append(froot, helpers.IntToString(item.BookStock))
	//		froot = append(froot, helpers.IntToString(item.PhysicalStock))
	//		datas = append(datas, froot)
	//}
	//title := "laporan-stock-opname"
	//headings := []string{"Produk", "Kategori", "Minimum Stok", "Stok Buku", "Stok Fisik"}
	//footer := map[string]float64{}
	//resultPdf, err := helpers.GeneratePdf(ctx, title, headings, datas, footer)

	//return helpers.Response("OK", "Sukses Export PDF", resultPdf), err
	return res, err
}
