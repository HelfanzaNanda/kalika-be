package repository

import (
	"errors"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type (
	CategoryRepository interface{
		Create(ctx echo.Context, db *gorm.DB, category *domain.Category) (domain.Category, error)
		Update(ctx echo.Context, db *gorm.DB, category *domain.Category) (domain.Category, error)
		Delete(ctx echo.Context, db *gorm.DB, category *domain.Category) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Category, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Category, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.CategoryDatatable, int64, int64, error)
	}

	CategoryRepositoryImpl struct {

	}
)

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository CategoryRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, category *domain.Category) (domain.Category, error) {
	db.Create(&category)
	categoryRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(category.Id))
	return categoryRes, nil
}

func (repository CategoryRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, category *domain.Category) (domain.Category, error) {
	db.Where("id = ?", category.Id).Updates(&category)
	categoryRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(category.Id))
	return categoryRes, nil
}

func (repository CategoryRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, category *domain.Category) (bool, error) {
	results := db.Where("id = ?", category.Id).Delete(&category)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|category tidak ditemukan")
	}
	return true, nil
}

func (repository CategoryRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (categoryRes domain.Category, err error) {
	results := db.Where(key+" = ?", value).First(&categoryRes)
	if results.RowsAffected < 1 {
		return categoryRes, errors.New("NOT_FOUND|category tidak ditemukan")
	}
	return categoryRes, nil
}

func (repository CategoryRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (categoryRes []domain.Category, err error) {
	db.Find(&categoryRes)
	return categoryRes, nil
}

func (repository CategoryRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.CategoryDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("categories").
		Select("divisions.id division_id, divisions.name division_name, categories.id, categories.name, categories.active").
		Joins("left join divisions on divisions.id = categories.division_id")

	qry.Count(&totalData)
	if search != "" {
		qry.Where("(categories.id = ? OR categories.name LIKE ?)", search, "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("categories.id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}