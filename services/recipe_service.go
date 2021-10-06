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
	RecipeService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindByProductId(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
		Datatable(ctx echo.Context) (res web.Datatable, err error)
	}

	RecipeServiceImpl struct {
		RecipeRepository repository.RecipeRepository
		RecipeDetailRepository repository.RecipeDetailRepository
		db *gorm.DB
	}
)

func NewRecipeService(RecipeRepository repository.RecipeRepository, RecipeDetailRepository repository.RecipeDetailRepository, db *gorm.DB) RecipeService {
	return &RecipeServiceImpl{
		RecipeRepository: RecipeRepository,
		RecipeDetailRepository: RecipeDetailRepository,
		db: db,
	}
}

func (service *RecipeServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	message := "Sukses Menyimpan Data"

	recipeRepo := domain.Recipe{}
	recipeDetailRepo := []web.RecipeDetailGet{}

	o := new(web.RecipePost)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	total := 0.0

	for _, val := range o.RecipeDetail {
		total += val.Quantity * val.Price
	}

	o.Total = total
	o.TotalCogs = total
	o.OverheadPrice = total + (total * (o.OverheadPercentage/100))
	o.RecommendationPrice = total + (total * (o.RecommendationPercentage/100))

	if o.Id > 0 {
		message = "Sukses Memperbarui Data"
		recipeRepo, err = service.RecipeRepository.Update(ctx, tx, &o.Recipe)
	} else {
		recipeRepo, err = service.RecipeRepository.Create(ctx, tx, &o.Recipe)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}
	o.Recipe = recipeRepo

	if o.Id > 0 {
		service.RecipeDetailRepository.DeleteByRecipe(ctx, tx, o.Id)
		recipeDetailRepo, err = service.RecipeDetailRepository.Create(ctx, tx, o)
	} else {
		recipeDetailRepo, err = service.RecipeDetailRepository.Create(ctx, tx, o)
	}
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}
	o.RecipeDetail = recipeDetailRepo

	return helpers.Response("CREATED", message, o), err
}

func (service *RecipeServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Recipe)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	recipeRepo, err := service.RecipeRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", recipeRepo), err
}

func (service *RecipeServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Recipe)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.RecipeRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service *RecipeServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	recipe := web.RecipePost{}
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	recipeRepo, err := service.RecipeRepository.FindById(ctx, tx, "id", helpers.IntToString(id))
	recipeDetailSearch := make(map[string][]string)
	recipeDetailSearch["recipe_id"] = append(recipeDetailSearch["recipe_id"], helpers.IntToString(recipeRepo.Id))
	recipeDetailRepo, err := service.RecipeDetailRepository.FindAll(ctx, tx, recipeDetailSearch)

	recipe.Recipe = recipeRepo
	recipe.RecipeDetail = recipeDetailRepo

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", recipe), err
}

func (service *RecipeServiceImpl) FindByProductId(ctx echo.Context, id int) (res web.Response, err error) {
	recipe := web.RecipePost{}
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	recipeRepo, err := service.RecipeRepository.FindById(ctx, tx, "product_id", helpers.IntToString(id))
	recipeDetailSearch := make(map[string][]string)
	recipeDetailSearch["recipe_id"] = append(recipeDetailSearch["recipe_id"], helpers.IntToString(recipeRepo.Id))
	recipeDetailRepo, err := service.RecipeDetailRepository.FindAll(ctx, tx, recipeDetailSearch)

	recipe.Recipe = recipeRepo
	recipe.RecipeDetail = recipeDetailRepo

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", recipe), err
}

func (service *RecipeServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	recipeRepo, err := service.RecipeRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", recipeRepo), err
}

func (service *RecipeServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params,_ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))

	recipeRepo, totalData, totalFiltered, _ := service.RecipeRepository.Datatable(ctx, tx, draw, limit, start, search)

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

