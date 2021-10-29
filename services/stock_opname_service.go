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
	StockOpnameService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
		Datatable(ctx echo.Context) (res web.Datatable, err error)
		GeneratePdf(ctx echo.Context, stockOpnameId int) (web.Response, error)
	}

	StockOpnameServiceImpl struct {
		StockOpnameRepository repository.StockOpnameRepository
		StockOpnameDetailRepository repository.StockOpnameDetailRepository
		PaymentRepository repository.PaymentRepository
		StoreConsignmentRepository repository.StoreConsignmentRepository
		db *gorm.DB
	}
)

func NewStockOpnameService(
	StockOpnameRepository repository.StockOpnameRepository,
	StockOpnameDetailRepository repository.StockOpnameDetailRepository,
	db *gorm.DB) StockOpnameService {
	return &StockOpnameServiceImpl{
		StockOpnameRepository: StockOpnameRepository,
		StockOpnameDetailRepository: StockOpnameDetailRepository,
		db: db,
	}
}

func (service *StockOpnameServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	message := "Sukses Menyimpan Data"
	o := new(web.StockOpnamePost)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	stockOpnameRepo := web.StockOpnameGet{}
	stockOpnameDetailRepo := []web.StockOpnameDetailGet{}

	o.StockOpname.CreatedBy = helpers.StringToInt(ctx.Get("userInfo").(map[string]interface{})["id"].(string))
	if o.Id > 0 {
		message = "Sukses Memperbarui Data"
		stockOpnameRepo, err = service.StockOpnameRepository.Update(ctx, tx, &o.StockOpname)
	} else {
		o.StockOpname.Number = "SC"+helpers.IntToString(int(time.Now().Unix()))
		o.StockOpname.Status = "submission"
		stockOpnameRepo, err = service.StockOpnameRepository.Create(ctx, tx, &o.StockOpname)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	//Assign Parent ID to Detail
	for key, _ := range o.StockOpnameDetail {
		o.StockOpnameDetail[key].StockOpnameId = stockOpnameRepo.Id
	}

	if o.Id > 0 {
		service.StockOpnameDetailRepository.DeleteByStockOpname(ctx, tx, o.Id)
		stockOpnameDetailRepo, err = service.StockOpnameDetailRepository.Create(ctx, tx, o.StockOpnameDetail)
	} else {
		stockOpnameDetailRepo, err = service.StockOpnameDetailRepository.Create(ctx, tx, o.StockOpnameDetail)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}
	stockOpnameRepo.StockOpnameDetail = stockOpnameDetailRepo

	return helpers.Response("CREATED", message, stockOpnameRepo), err
}

func (service *StockOpnameServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.StockOpname)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	stockOpnameRepo, err := service.StockOpnameRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", stockOpnameRepo), err
}

func (service *StockOpnameServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.StockOpname)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.StockOpnameRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service *StockOpnameServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	stockOpname := web.StockOpnameGet{}
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	stockOpnameRepo, err := service.StockOpnameRepository.FindById(ctx, tx, "id", helpers.IntToString(id))
	stockOpnameDetailSearch := make(map[string][]string)
	stockOpnameDetailSearch["stock_opname_id"] = append(stockOpnameDetailSearch["stock_opname_id"], helpers.IntToString(stockOpnameRepo.Id))
	stockOpnameDetailRepo, err := service.StockOpnameDetailRepository.FindAll(ctx, tx, stockOpnameDetailSearch)

	stockOpname = stockOpnameRepo
	stockOpname.StockOpnameDetail = stockOpnameDetailRepo

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", stockOpname), err
}

func (service *StockOpnameServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	stockOpnameRepo, err := service.StockOpnameRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", stockOpnameRepo), err
}

func (service *StockOpnameServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params,_ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))

	recipeRepo, totalData, totalFiltered, _ := service.StockOpnameRepository.Datatable(ctx, tx, draw, limit, start, search)

	data := make([]interface{}, 0)
	for _, v := range recipeRepo {
		v.Action = `<div class="flex">`
		v.Action += `<button type="button" class="btn-edit flex " id="edit-data" data-id=`+helpers.IntToString(v.Id)+`> <i data-feather="check-square" class="w-4 h-4 mr-1"></i> Edit </button>`
		v.Action += `<button type="button" class="btn-delete flex mx-3 text-theme-6" id="delete-data" data-id=`+helpers.IntToString(v.Id)+`> <i data-feather="trash-2" class="w-4 h-4 mr-1"></i> Delete </button>`
		v.Action += `<button type="button" class="btn-pdf flex text-theme-12" id="pdf-data" data-id=`+helpers.IntToString(v.Id)+`> <i data-feather="file-text" class="w-4 h-4 mr-1"></i> Cetak </button>`
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

func (service StockOpnameServiceImpl) GeneratePdf(ctx echo.Context, stockOpnameId int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	stockOpnameRepo, err := service.StockOpnameDetailRepository.Pdf(ctx, tx, stockOpnameId)
	var datas [][]string
	for _, item := range stockOpnameRepo {
		froot := []string{}
			froot = append(froot, item.ProductName)
			froot = append(froot, item.CategoryName)
			froot = append(froot, helpers.IntToString(item.MinimumStock))
			froot = append(froot, helpers.IntToString(item.BookStock))
			froot = append(froot, helpers.IntToString(item.PhysicalStock))
			datas = append(datas, froot)
	}
	title := "laporan-stock-opname"
	headings := []string{"Produk", "Kategori", "Minimum Stok", "Stok Buku", "Stok Fisik"}
	footer := map[string]float64{}
	resultPdf, err := helpers.GeneratePdf(ctx, title, headings, datas, footer)

	return helpers.Response("OK", "Sukses Export PDF", resultPdf), err
}
