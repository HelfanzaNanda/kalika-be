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
	CakeVariantRepository interface{
		Create(ctx echo.Context, db *gorm.DB, cakeVariant *domain.CakeVariant) (domain.CakeVariant, error)
		Update(ctx echo.Context, db *gorm.DB, cakeVariant *domain.CakeVariant) (domain.CakeVariant, error)
		Delete(ctx echo.Context, db *gorm.DB, cakeVariant *domain.CakeVariant) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.CakeVariant, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.CakeVariant, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.CakeVariantDatatable, int64, int64, error)
	}

	CakeVariantRepositoryImpl struct {

	}
)

func NewCakeVariantRepository() CakeVariantRepository {
	return &CakeVariantRepositoryImpl{}
}

func (repository CakeVariantRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, cakeVariant *domain.CakeVariant) (domain.CakeVariant, error) {
	db.Create(&cakeVariant)
	cakeVariantRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(cakeVariant.Id))
	return cakeVariantRes, nil
}

func (repository CakeVariantRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, cakeVariant *domain.CakeVariant) (domain.CakeVariant, error) {
	db.Where("id = ?", cakeVariant.Id).Updates(&cakeVariant)
	cakeVariantRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(cakeVariant.Id))
	return cakeVariantRes, nil
}

func (repository CakeVariantRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, cakeVariant *domain.CakeVariant) (bool, error) {
	results := db.Where("id = ?", cakeVariant.Id).Delete(&cakeVariant)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|cakeVariant tidak ditemukan")
	}
	return true, nil
}

func (repository CakeVariantRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (cakeVariantRes domain.CakeVariant, err error) {
	results := db.Where(key+" = ?", value).First(&cakeVariantRes)
	if results.RowsAffected < 1 {
		return cakeVariantRes, errors.New("NOT_FOUND|cakeVariant tidak ditemukan")
	}
	return cakeVariantRes, nil
}

func (repository CakeVariantRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (cakeVariantRes []domain.CakeVariant, err error) {
	db.Find(&cakeVariantRes)
	return cakeVariantRes, nil
}

func (repository CakeVariantRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.CakeVariantDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("cake_variants")
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(id = ? OR name LIKE ?)", search, "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}
