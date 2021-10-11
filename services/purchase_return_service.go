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
	PurchaseReturnService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
		ReportDatatable(ctx echo.Context) (res web.Datatable, err error)
		Datatable(ctx echo.Context) (res web.Datatable, err error)
		GeneratePdf(ctx echo.Context) (web.Response, error)
	}

	PurchaseReturnServiceImpl struct {
		PurchaseReturnRepository repository.PurchaseReturnRepository
		PurchaseReturnDetailRepository repository.PurchaseReturnDetailRepository
		db *gorm.DB
	}
)

func NewPurchaseReturnService(PurchaseReturnRepository repository.PurchaseReturnRepository, PurchaseReturnDetailRepository repository.PurchaseReturnDetailRepository, db *gorm.DB) PurchaseReturnService {
	return &PurchaseReturnServiceImpl{
		PurchaseReturnRepository: PurchaseReturnRepository,
		PurchaseReturnDetailRepository: PurchaseReturnDetailRepository,
		db: db,
	}
}

func (service *PurchaseReturnServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(web.PurchaseReturnPost)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)
	purchaseReturnRepo, err := service.PurchaseReturnRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "create purchase return error", nil), err
	}
	o.PurchaseReturn = purchaseReturnRepo

	purchaseReturnDetailRepo, err := service.PurchaseReturnDetailRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "create purchase return detail error", nil), err
	}
	
	o.PurchaseReturnDetails = purchaseReturnDetailRepo
	return helpers.Response("OK", "Sukses Menyimpan Data", o), err
}

func (service PurchaseReturnServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(web.PurchaseReturnPost)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)
	purchaseReturnRepo, err := service.PurchaseReturnRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "create purchase return error", nil), err
	}
	o.PurchaseReturn = purchaseReturnRepo

	_, err = service.PurchaseReturnDetailRepository.DeleteByPurchaseReturnId(ctx, tx, id)
	if err != nil {
		return helpers.Response(err.Error(), "delete purchase return detail error", nil), err
	}
	
	o.Id = id
	purchaseReturnDetailRepo, err := service.PurchaseReturnDetailRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "create purchase return detail error", nil), err
	}
	
	o.PurchaseReturnDetails = purchaseReturnDetailRepo
	return helpers.Response("OK", "Sukses Mengubah Data", o), err
}

func (service PurchaseReturnServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.PurchaseReturn)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.PurchaseReturnRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "delete purchase return error", nil), err
	}
	_, err = service.PurchaseReturnDetailRepository.DeleteByPurchaseReturnId(ctx, tx, id)
	if err != nil {
		return helpers.Response(err.Error(), "delete purchase return detail error", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service PurchaseReturnServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	purxhaseReturn := web.PurchaseReturnPost{}
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	purchaseReturnRepo, err := service.PurchaseReturnRepository.FindById(ctx, tx, "id", helpers.IntToString(id))
	
	purchaseReturnDetailRepo, err := service.PurchaseReturnDetailRepository.FindByPurchaseReturnId(ctx, tx, id)

	purxhaseReturn.PurchaseReturn = purchaseReturnRepo
	purxhaseReturn.Date = purchaseReturnRepo.Date.String()
	purxhaseReturn.PurchaseReturnDetails = purchaseReturnDetailRepo

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", purxhaseReturn), err
}

func (service PurchaseReturnServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	purchaseReturnRepo, err := service.PurchaseReturnRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", purchaseReturnRepo), err
}

func (service *PurchaseReturnServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params,_ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))

	purchaseReturnRepo, totalData, totalFiltered, _ := service.PurchaseReturnRepository.Datatable(ctx, tx, draw, limit, start, search)
	// if err != nil {
	// 	return helpers.Response(err.Error(), "", nil), err
	// }

	data := make([]interface{}, 0)
	for _, v := range purchaseReturnRepo {
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

func (service *PurchaseReturnServiceImpl) ReportDatatable(ctx echo.Context) (res web.Datatable, err error) {
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
	purchaseReturnRepo, totalData, totalFiltered, _ := service.PurchaseReturnRepository.ReportDatatable(ctx, tx, draw, limit, start, search, filter)
	// if err != nil {
	// 	return helpers.Response(err.Error(), "", nil), err
	// }

	data := make([]interface{}, 0)
	for _, v := range purchaseReturnRepo {
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

func (service PurchaseReturnServiceImpl) GeneratePdf(ctx echo.Context) (res web.Response, err error) {
	o := new(web.DateRange)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)
	
	purchaseReturnRepo, err := service.PurchaseReturnRepository.FindByCreatedAt(ctx, tx, o)
	var datas [][]string
	for _, item := range purchaseReturnRepo {
		froot := []string{}
		froot = append(froot, item.Date.String())
		froot = append(froot, item.Number)
		
		datas = append(datas, froot)
	}
	title := "laporan-return-pembelian"
	headings := []string{"Date", "Number"}
	resultPdf, err := helpers.GeneratePdf(ctx, title, headings, datas)
	
	return helpers.Response("OK", "Sukses Export PDF", resultPdf), err
}