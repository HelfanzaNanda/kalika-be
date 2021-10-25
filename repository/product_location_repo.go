package repository

import (
	"errors"
	"fmt"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type (
	ProductLocationRepository interface{
		Create(ctx echo.Context, db *gorm.DB, product []domain.ProductLocation) ([]domain.ProductLocation, error)
		Update(ctx echo.Context, db *gorm.DB, product *domain.ProductLocation) (web.ProductLocationGet, error)
		Delete(ctx echo.Context, db *gorm.DB, product *domain.ProductLocation) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (web.ProductLocationGet, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]web.ProductLocationGet, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.ProductDatatable, int64, int64, error)
		DeleteByProduct(ctx echo.Context, db *gorm.DB, model string, productId int) (bool, error)
		StockDeduction(ctx echo.Context, db *gorm.DB, params []map[string]interface{}) (bool, error)
		StockAddition(ctx echo.Context, db *gorm.DB, params []map[string]interface{}) (bool, error)
		CheckStockDataTable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string, filter map[string]string) ([]web.CheckStockDataTable, int64, int64, error)
		CheckStockPdf(ctx echo.Context, db *gorm.DB, filter *web.CheckStockFilter) ([]web.CheckStockGet, error)
	}

	ProductLocationRepositoryImpl struct {

	}
)

func NewProductLocationRepository() ProductLocationRepository {
	return &ProductLocationRepositoryImpl{}
}

func (r *ProductLocationRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, product []domain.ProductLocation) (res []domain.ProductLocation, err error) {
	for _, val := range product {
		db.Create(&val)
		res = append(res, val)
	}
	return res, nil
}

func (r *ProductLocationRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, product *domain.ProductLocation) (web.ProductLocationGet, error) {
	db.Where("id = ?", product.Id).Updates(&product)
	productLocationRes,_ := r.FindById(ctx, db, "id", helpers.IntToString(product.Id))
	return productLocationRes, nil
}

func (r *ProductLocationRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, product *domain.ProductLocation) (bool, error) {
	results := db.Where("id = ?", product.Id).Delete(&product)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|product tidak ditemukan")
	}
	return true, nil
}

func (r *ProductLocationRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (productLocationRes web.ProductLocationGet, err error) {
	results := db.Where(key+" = ?", value).First(&productLocationRes)
	if results.RowsAffected < 1 {
		return productLocationRes, errors.New("NOT_FOUND|product tidak ditemukan")
	}
	return productLocationRes, nil
}

func (r *ProductLocationRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (productLocationRes []web.ProductLocationGet, err error) {
	productLocations := []domain.ProductLocation{}
	product := domain.Product{}
	categoryProduct := domain.Category{}
	divisionProduct := domain.Division{}
	store := domain.Store{}

	qry := db.Table("product_locations").Select("product_locations.*")
	for k, v := range ctx.QueryParams() {
		if v[0] != "" && k != "id" {
			qry = qry.Where(k+" = ?", v[0])
		}
	}
	qry.Scan(&productLocations)

	for _, val := range productLocations {
		result := web.ProductLocationGet{}
		result.ProductLocation = val
		db.Model(product).Where("id", val.ProductId).First(&result.Product)
		db.Model(categoryProduct).Where("id", result.Product.CategoryId).First(&result.Product.Category)
		db.Model(divisionProduct).Where("id", result.Product.DivisionId).First(&result.Product.Division)
		db.Model(store).Where("id", val.StoreId).First(&result.Store)
		productLocationRes = append(productLocationRes, result)
	}
	return productLocationRes, nil
}


func (r *ProductLocationRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.ProductDatatable, totalData int64, totalFiltered int64, err error) {
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

func (r *ProductLocationRepositoryImpl) DeleteByProduct(ctx echo.Context, db *gorm.DB, model string, productId int) (bool, error) {
	db.Where("product_id = ? AND model = ?", productId, model).Delete(domain.ProductLocation{})
	return true, nil
}

func (r *ProductLocationRepositoryImpl) StockDeduction(ctx echo.Context, db *gorm.DB, params []map[string]interface{}) (bool, error) {
	var model string
	var productId int
	var quantity int
	var storeId int
	for _, val := range params {
		model = val["model"].(string)
		productId = helpers.StringToInt(val["product_id"].(string))
		quantity = helpers.StringToInt(val["quantity"].(string))
		storeId = helpers.StringToInt(val["store_id"].(string))

		update := db.Model(&domain.ProductLocation{}).Where("model = ? AND product_id = ? AND store_id = ?", model, productId, storeId).Update("quantity", gorm.Expr("quantity - ?", quantity))
		if update.RowsAffected < 1 {
			fmt.Println(fmt.Sprint("No stock detected on model = ",model, " AND product_id = ", productId, " AND store_id = ", storeId))
		}
	}
	return false, nil
}

func (r *ProductLocationRepositoryImpl) StockAddition(ctx echo.Context, db *gorm.DB, params []map[string]interface{}) (bool, error) {
	var model string
	var productId int
	var quantity int
	var storeId int
	for _, val := range params {
		model = val["model"].(string)
		productId = helpers.StringToInt(val["product_id"].(string))
		quantity = helpers.StringToInt(val["quantity"].(string))
		storeId = helpers.StringToInt(val["store_id"].(string))

		update := db.Model(&domain.ProductLocation{}).Where("model = ? AND product_id = ? AND store_id = ?", model, productId, storeId).Update("quantity", gorm.Expr("quantity + ?", quantity))
		if update.RowsAffected < 1 {
			fmt.Println(fmt.Sprint("No stock detected on model = ",model, " AND product_id = ", productId, " AND store_id = ", storeId))
		}
	}
	return false, nil
}

func (repository ProductLocationRepositoryImpl) CheckStockDataTable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string, filter map[string]string) (res []web.CheckStockDataTable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("product_locations")
	qry.Select(`
		product_locations.quantity qty,
		products.name product_name, products.stock_minimum minimum_stock,
		divisions.name division_name, categories.name category_name
	`)
	qry.Joins(`
		JOIN products ON products.id = product_locations.product_id
		JOIN stores ON stores.id = product_locations.store_id
		JOIN categories ON categories.id = products.category_id
		JOIN divisions ON divisions.id = categories.division_id
	`)
	qry.Where("(product_locations.model = ?)", "Product")
	qry.Count(&totalData)

	if filter["division_id"] != "" {
		qry.Where("(categories.division_id = ?)", filter["division_id"])
	}
	if filter["store_id"] != "" {
		qry.Where("(product_locations.store_id = ?)", filter["store_id"])
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("categories.division_id, products.category_id")
	qry.Find(&res)
	return res, totalData, totalFiltered, nil
}

func (repository ProductLocationRepositoryImpl) CheckStockPdf(ctx echo.Context, db *gorm.DB, filter *web.CheckStockFilter) (res []web.CheckStockGet, err error) {
	qry := db.Table("product_locations")
	qry.Select(`
		product_locations.quantity qty,
		products.name product_name, products.stock_minimum minimum_stock,
		divisions.name division_name, categories.name category_name
	`)
	qry.Joins(`
		left JOIN products ON products.id = product_locations.product_id
		left JOIN stores ON stores.id = product_locations.store_id
		left JOIN categories ON categories.id = products.category_id
		JOIN divisions ON divisions.id = categories.division_id
	`)
	qry.Where("(product_locations.model = ?)", "Product")
	
	if filter.StoreId != 0 {
		qry.Where("(product_locations.store_id = ?)", filter.StoreId)
	}
	if filter.DivisionId != 0 {
		qry.Where("(categories.division_id = ?)", filter.DivisionId)
	}
	qry.Order("categories.division_id, products.category_id")
	qry.Find(&res)
	return res, nil
}