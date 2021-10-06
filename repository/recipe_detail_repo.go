package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"
)

type (
	RecipeDetailRepository interface{
		Create(ctx echo.Context, db *gorm.DB, recipe *web.RecipePost) ([]web.RecipeDetailGet, error)
		Update(ctx echo.Context, db *gorm.DB, recipeDetail *domain.RecipeDetail) (domain.RecipeDetail, error)
		Delete(ctx echo.Context, db *gorm.DB, recipeDetail *domain.RecipeDetail) (bool, error)
		DeleteByRecipe(ctx echo.Context, db *gorm.DB, recipeId int) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.RecipeDetail, error)
		FindAll(ctx echo.Context, db *gorm.DB, params map[string][]string) ([]web.RecipeDetailGet, error)
	}

	RecipeDetailRepositoryImpl struct {

	}
)

func NewRecipeDetailRepository() RecipeDetailRepository {
	return &RecipeDetailRepositoryImpl{}
}

func (repository RecipeDetailRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, recipe *web.RecipePost) (res []web.RecipeDetailGet, err error) {
	for _, val := range recipe.RecipeDetail {
		val.RecipeId = recipe.Id
		val.Total = val.Quantity * val.Price
		db.Table("recipe_details").Select("recipe_id", "raw_material_id", "price", "unit_id", "quantity", "total").Create(&val)
		res = append(res, val)
	}

	return res, nil
}

func (repository RecipeDetailRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, recipeDetail *domain.RecipeDetail) (domain.RecipeDetail, error) {
	db.Where("id = ?", recipeDetail.Id).Updates(&recipeDetail)
	recipeDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(recipeDetail.Id))
	return recipeDetailRes, nil
}

func (repository RecipeDetailRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, recipeDetail *domain.RecipeDetail) (bool, error) {
	results := db.Where("id = ?", recipeDetail.Id).Delete(&recipeDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|recipeDetail tidak ditemukan")
	}
	return true, nil
}

func (repository RecipeDetailRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (recipeDetailRes domain.RecipeDetail, err error) {
	results := db.Where(key+" = ?", value).First(&recipeDetailRes)
	if results.RowsAffected < 1 {
		return recipeDetailRes, errors.New("NOT_FOUND|recipeDetail tidak ditemukan")
	}
	return recipeDetailRes, nil
}

func (repository RecipeDetailRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB, params map[string][]string) (recipeDetailRes []web.RecipeDetailGet, err error) {
	results := db.Table("recipe_details").Preload("RawMaterial")
	for k, v := range params {
		if v[0] != "" && k != "id" {
			results = results.Where(k+" = ?", v[0])
		}
	}

	results.Find(&recipeDetailRes)
	return recipeDetailRes, nil
}

func (repository RecipeDetailRepositoryImpl) DeleteByRecipe(ctx echo.Context, db *gorm.DB, recipeId int) (bool, error) {
	recipeDetail := domain.RecipeDetail{}
	results := db.Where("recipe_id = ?", recipeId).Delete(&recipeDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|recipeDetail tidak ditemukan")
	}
	return true, nil
}