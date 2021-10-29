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
	StoreMutationService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
		Datatable(ctx echo.Context) (res web.Datatable, err error)
		GeneratePdf(ctx echo.Context, storeMutationId int) (web.Response, error)
	}

	StoreMutationServiceImpl struct {
		StoreMutationRepository repository.StoreMutationRepository
		StoreMutationDetailRepository repository.StoreMutationDetailRepository
		PaymentRepository repository.PaymentRepository
		StoreConsignmentRepository repository.StoreConsignmentRepository
		db *gorm.DB
	}
)

func NewStoreMutationService(
	StoreMutationRepository repository.StoreMutationRepository,
	StoreMutationDetailRepository repository.StoreMutationDetailRepository,
	db *gorm.DB) StoreMutationService {
	return &StoreMutationServiceImpl{
		StoreMutationRepository: StoreMutationRepository,
		StoreMutationDetailRepository: StoreMutationDetailRepository,
		db: db,
	}
}

func (service *StoreMutationServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	message := "Sukses Menyimpan Data"
	o := new(web.StoreMutationPost)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	storeMutationRepo := web.StoreMutationGet{}
	storeMutationDetailRepo := []web.StoreMutationDetailGet{}

	o.StoreMutation.CreatedBy = helpers.StringToInt(ctx.Get("userInfo").(map[string]interface{})["id"].(string))
	if o.Id > 0 {
		message = "Sukses Memperbarui Data"
		storeMutationRepo, err = service.StoreMutationRepository.Update(ctx, tx, &o.StoreMutation)
	} else {
		o.StoreMutation.Number = "SM"+helpers.IntToString(int(time.Now().Unix()))
		o.StoreMutation.Status = "submission"
		storeMutationRepo, err = service.StoreMutationRepository.Create(ctx, tx, &o.StoreMutation)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	for key, _ := range o.StoreMutationDetail {
		o.StoreMutationDetail[key].StoreMutationId = storeMutationRepo.Id
	}

	if o.Id > 0 {
		service.StoreMutationDetailRepository.DeleteByStoreMutation(ctx, tx, o.Id)
		storeMutationDetailRepo, err = service.StoreMutationDetailRepository.Create(ctx, tx, o.StoreMutationDetail)
	} else {
		storeMutationDetailRepo, err = service.StoreMutationDetailRepository.Create(ctx, tx, o.StoreMutationDetail)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}
	storeMutationRepo.StoreMutationDetail = storeMutationDetailRepo

	return helpers.Response("CREATED", message, storeMutationRepo), err
}

func (service *StoreMutationServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.StoreMutation)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	storeMutationRepo, err := service.StoreMutationRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", storeMutationRepo), err
}

func (service *StoreMutationServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.StoreMutation)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.StoreMutationRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service *StoreMutationServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	storeMutation := web.StoreMutationGet{}
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	storeMutationRepo, err := service.StoreMutationRepository.FindById(ctx, tx, "id", helpers.IntToString(id))
	storeMutationDetailSearch := make(map[string][]string)
	storeMutationDetailSearch["store_mutation_id"] = append(storeMutationDetailSearch["store_mutation_id"], helpers.IntToString(storeMutationRepo.Id))
	storeMutationDetailRepo, err := service.StoreMutationDetailRepository.FindAll(ctx, tx, storeMutationDetailSearch)

	storeMutation = storeMutationRepo
	storeMutation.StoreMutationDetail = storeMutationDetailRepo

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", storeMutation), err
}

func (service *StoreMutationServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	storeMutationRepo, err := service.StoreMutationRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", storeMutationRepo), err
}

func (service *StoreMutationServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params,_ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))

	recipeRepo, totalData, totalFiltered, _ := service.StoreMutationRepository.Datatable(ctx, tx, draw, limit, start, search)

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

func (service StoreMutationServiceImpl) GeneratePdf(ctx echo.Context, storeMutationId int) (res web.Response, err error) {
	//tx := service.db.Begin()
	//defer helpers.CommitOrRollback(tx)
	//
	//storeMutationRepo, err := service.StoreMutationDetailRepository.Pdf(ctx, tx, storeMutationId)
	//var datas [][]string
	//for _, item := range storeMutationRepo {
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
	//
	//return helpers.Response("OK", "Sukses Export PDF", resultPdf), err

	return res, err
}
