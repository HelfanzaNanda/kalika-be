package services

import (
	"strings"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/web"
	"kalika-be/repository"
)

type (
	CheckStockService interface {
		Datatable(ctx echo.Context) (res web.Datatable, err error)
		GeneratePdf(ctx echo.Context) (web.Response, error)
	}

	CheckStockServiceImpl struct {
		ProductLocationRepository repository.ProductLocationRepository
		db *gorm.DB
	}
)

func NewCheckStockService(ProductLocationRepository repository.ProductLocationRepository, db *gorm.DB) CheckStockService {
	return &CheckStockServiceImpl{
		ProductLocationRepository: ProductLocationRepository,
		db: db,
	}
}

func (service *CheckStockServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params, _ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))
	filter := make(map[string]string)
	filter["division_id"] = strings.TrimSpace(params.Get("division_id"))
	filter["store_id"] = strings.TrimSpace(params.Get("store_id"))
	customOrderRepo, totalData, totalFiltered, _ := service.ProductLocationRepository.CheckStockDataTable(ctx, tx, draw, limit, start, search, filter)

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

func (service CheckStockServiceImpl) GeneratePdf(ctx echo.Context) (res web.Response, err error) {
	o := new(web.CheckStockFilter)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	productLocationRepo, err := service.ProductLocationRepository.CheckStockPdf(ctx, tx, o)
	var datas [][]string
	for _, item := range productLocationRepo {
		froot := []string{}
			froot = append(froot, item.ProductName)
			froot = append(froot, helpers.IntToString(item.Qty))
			froot = append(froot, helpers.IntToString(item.MinimumStock))
			froot = append(froot, item.DivisionName)
			froot = append(froot, item.CategoryName)
			datas = append(datas, froot)
	}
	title := "laporan_check_stok"
	headings := []string{"Produk", "Stok", "Minimum Stok", "Divisi", "Kategori"}
	footer := map[string]float64{}
	resultPdf, err := helpers.GeneratePdf(ctx, title, headings, datas, footer, "", "")

	return helpers.Response("OK", "Sukses Export PDF", resultPdf), err
}
