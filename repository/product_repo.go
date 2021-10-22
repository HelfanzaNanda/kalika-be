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
	ProductRepository interface{
		Create(ctx echo.Context, db *gorm.DB, product *web.ProductPosPost) (domain.Product, error)
		Update(ctx echo.Context, db *gorm.DB, product *web.ProductPosPost) (domain.Product, error)
		Delete(ctx echo.Context, db *gorm.DB, product *domain.Product) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Product, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Product, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.ProductDatatable, int64, int64, error)
	}

	ProductRepositoryImpl struct {

	}
)

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository ProductRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, product *web.ProductPosPost) (domain.Product, error) {
	m := domain.Product{}
	m.Name = product.Name
	m.StockMinimum = product.StockMinimum
	m.ProductionMinimum = product.ProductionMinimum
	m.DivisionId = product.DivisionId
	m.CakeTypeId = product.CakeTypeId
	m.CategoryId = product.CategoryId
	m.Active = product.Active
	m.IsCustomPrice = product.IsCustomPrice
	m.IsCustomProduct = product.IsCustomProduct
	db.Create(&m)
	productRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(m.Id))
	return productRes, nil
}

func (repository ProductRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, product *web.ProductPosPost) (domain.Product, error) {
	db.Where("id = ?", product.Id).Updates(&product.Product)
	productRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(product.Id))
	return productRes, nil
}

func (repository ProductRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, product *domain.Product) (bool, error) {
	results := db.Where("id = ?", product.Id).Delete(&product)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|product tidak ditemukan")
	}
	return true, nil
}

func (repository ProductRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (productRes domain.Product, err error) {
	results := db.Where(key+" = ?", value).First(&productRes)
	if results.RowsAffected < 1 {
		return productRes, errors.New("NOT_FOUND|product tidak ditemukan")
	}
	return productRes, nil
}

func (repository ProductRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (productRes []domain.Product, err error) {
	qry := db.Table("products").Select("products.*")
	for k, v := range ctx.QueryParams() {
		if k == "name" {
			qry = qry.Where(k+" LIKE ?", "%"+v[0]+"%")
		} else if v[0] != "" && k != "id" {
			qry = qry.Where(k+" = ?", v[0])
		}
	}
	qry.Scan(&productRes)

	for key, val := range productRes {
		db.Table("product_prices").Where("product_id = ?", val.Id).Scan(&productRes[key].ProductPrice)
	}

	return productRes, nil
}


func (repository ProductRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.ProductDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("products").
	Select(`
		products.id, products.name, products.stock_minimum, products.production_minimum, 
		products.active, products.is_custom_price, products.is_custom_product,
		divisions.id division_id, divisions.name division_name,
		categories.id category_id, categories.name category_name,
		cake_variants.id cake_variant_id, cake_variants.name cake_variant_name,
		cake_types.id cake_type_id, cake_types.name cake_type_name
	`).
	Joins(`
		left join categories on categories.id = products.category_id
		left join divisions on divisions.id = categories.division_id
		left join cake_variants on cake_variants.id = products.cake_variant_id
		left join cake_types on cake_types.id = products.cake_type_id
	`)
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(products.id = ? OR products.name LIKE ?)", search, "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("products.id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}
