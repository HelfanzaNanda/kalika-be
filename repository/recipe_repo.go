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
	RecipeRepository interface{
		Create(ctx echo.Context, db *gorm.DB, role *domain.Recipe) (domain.Recipe, error)
		Update(ctx echo.Context, db *gorm.DB, role *domain.Recipe) (domain.Recipe, error)
		Delete(ctx echo.Context, db *gorm.DB, role *domain.Recipe) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Recipe, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Recipe, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.RecipeDatatable, int64, int64, error)
	}

	RecipeRepositoryImpl struct {

	}
)

func NewRecipeRepository() RecipeRepository {
	return &RecipeRepositoryImpl{}
}

func (repository *RecipeRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, role *domain.Recipe) (domain.Recipe, error) {
	db.Create(&role)
	recipeRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(role.Id))
	return recipeRes, nil
}

func (repository *RecipeRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, role *domain.Recipe) (domain.Recipe, error) {
	db.Where("id = ?", role.Id).Updates(&role)
	recipeRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(role.Id))
	return recipeRes, nil
}

func (repository *RecipeRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, role *domain.Recipe) (bool, error) {
	results := db.Where("id = ?", role.Id).Delete(&role)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|role tidak ditemukan")
	}
	return true, nil
}

func (repository *RecipeRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (recipeRes domain.Recipe, err error) {
	results := db.Where(key+" = ?", value).First(&recipeRes)
	if results.RowsAffected < 1 {
		return recipeRes, errors.New("NOT_FOUND|role tidak ditemukan")
	}
	return recipeRes, nil
}

func (repository *RecipeRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (recipeRes []domain.Recipe, err error) {
	db.Find(&recipeRes)
	return recipeRes, nil
}

func (repository *RecipeRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (recipeRes []web.RecipeDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("recipes")
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(products.name LIKE ?)", "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("recipes.id desc")
	qry.Preload("Product").Find(&recipeRes)
	return recipeRes, totalData, totalFiltered, nil
}

