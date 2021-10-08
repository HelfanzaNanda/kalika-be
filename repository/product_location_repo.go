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
	ProductLocationRepository interface{
		Create(ctx echo.Context, db *gorm.DB, product *web.ProductPosPost) (web.ProductPosPost, error)
		Update(ctx echo.Context, db *gorm.DB, product *domain.ProductLocation) (web.ProductLocationGet, error)
		Delete(ctx echo.Context, db *gorm.DB, product *domain.ProductLocation) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (web.ProductLocationGet, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.ProductLocation, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.ProductDatatable, int64, int64, error)
		DeleteByProduct(ctx echo.Context, db *gorm.DB, productId int) (bool, error)
	}

	ProductLocationRepositoryImpl struct {

	}
)

func NewProductLocationRepository() ProductLocationRepository {
	return &ProductLocationRepositoryImpl{}
}

func (repository ProductLocationRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, product *web.ProductPosPost) (res web.ProductPosPost, err error) {
	for _, val := range product.ProductLocations {
		val.ProductId = product.Id
		db.Create(&val)
		res.ProductLocations = append(res.ProductLocations, val)
	}
	return res, nil
}

func (repository ProductLocationRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, product *domain.ProductLocation) (web.ProductLocationGet, error) {
	db.Where("id = ?", product.Id).Updates(&product)
	productLocationRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(product.Id))
	return productLocationRes, nil
}

func (repository ProductLocationRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, product *domain.ProductLocation) (bool, error) {
	results := db.Where("id = ?", product.Id).Delete(&product)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|product tidak ditemukan")
	}
	return true, nil
}

func (repository ProductLocationRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (productLocationRes web.ProductLocationGet, err error) {
	results := db.Where(key+" = ?", value).First(&productLocationRes)
	if results.RowsAffected < 1 {
		return productLocationRes, errors.New("NOT_FOUND|product tidak ditemukan")
	}
	return productLocationRes, nil
}

func (repository ProductLocationRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (productLocationRes []domain.ProductLocation, err error) {
	qry := db.Table("product_locations").Select("product_locations.*")
	for k, v := range ctx.QueryParams() {
		if v[0] != "" && k != "id" {
			qry = qry.Where(k+" = ?", v[0])
		}
	}
	qry.Scan(&productLocationRes)
	return productLocationRes, nil
}


func (repository ProductLocationRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.ProductDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("product_locations").
	Select(`
		product_locations.id, product_locations.name, product_locations.stock_minimum, product_locations.production_minimum, 
		product_locations.active, product_locations.is_custom_price, product_locations.is_custom_product,
		divisions.id division_id, divisions.name division_name,
		categories.id category_id, categories.name category_name,
		cake_variants.id cake_variant_id, cake_variants.name cake_variant_name,
		cake_types.id cake_type_id, cake_types.name cake_type_name
	`).
	Joins(`
		left join divisions on divisions.id = product_locations.division_id
		left join categories on categories.id = product_locations.category_id
		left join cake_variants on cake_variants.id = product_locations.cake_variant_id
		left join cake_types on cake_types.id = product_locations.cake_type_id
	`)
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(product_locations.id = ? OR product_locations.name LIKE ?)", search, "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("product_locations.id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}

func (repository ProductLocationRepositoryImpl) DeleteByProduct(ctx echo.Context, db *gorm.DB, productId int) (bool, error) {
	db.Where("product_id = ?", productId).Delete(domain.ProductLocation{})
	return true, nil
}
