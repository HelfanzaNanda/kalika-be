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
	ProductionRequestDetailRepository interface{
		Create(echo.Context, *gorm.DB, []domain.ProductionRequestDetail) ([]web.ProductionRequestDetailGet, error)
		Update(echo.Context, *gorm.DB, *domain.ProductionRequestDetail) (domain.ProductionRequestDetail, error)
		Delete(echo.Context, *gorm.DB, *domain.ProductionRequestDetail) (bool, error)
		DeleteByProductionRequest(echo.Context, *gorm.DB, int) (bool, error)
		FindById(echo.Context, *gorm.DB, string, string) (domain.ProductionRequestDetail, error)
		FindAll(echo.Context, *gorm.DB, map[string][]string) ([]web.ProductionRequestDetailGet, error)
		Pdf(echo.Context, *gorm.DB, int) ([]web.ProductionRequestDetailReportGet, error)
	}

	ProductionRequestDetailRepositoryImpl struct {

	}
)

func NewProductionRequestDetailRepository() ProductionRequestDetailRepository {
	return &ProductionRequestDetailRepositoryImpl{}
}

func (repository *ProductionRequestDetailRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, productionRequestDetail []domain.ProductionRequestDetail) (res []web.ProductionRequestDetailGet, err error) {
	for _, val := range productionRequestDetail {
		db.Create(&val)
	}

	detailSearch := make(map[string][]string)
	detailSearch["production_request_id"] = append(detailSearch["production_request_id"], helpers.IntToString(productionRequestDetail[0].ProductionRequestId))

	res, _ = repository.FindAll(ctx, db, detailSearch)

	return res, nil
}

func (repository *ProductionRequestDetailRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, productionRequestDetail *domain.ProductionRequestDetail) (domain.ProductionRequestDetail, error) {
	db.Where("id = ?", productionRequestDetail.Id).Updates(&productionRequestDetail)
	productionRequestDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(productionRequestDetail.Id))
	return productionRequestDetailRes, nil
}

func (repository *ProductionRequestDetailRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, productionRequestDetail *domain.ProductionRequestDetail) (bool, error) {
	results := db.Where("id = ?", productionRequestDetail.Id).Delete(&productionRequestDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|productionRequestDetail tidak ditemukan")
	}
	return true, nil
}

func (repository *ProductionRequestDetailRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (productionRequestDetailRes domain.ProductionRequestDetail, err error) {
	results := db.Where(key+" = ?", value).First(&productionRequestDetailRes)
	if results.RowsAffected < 1 {
		return productionRequestDetailRes, errors.New("NOT_FOUND|productionRequestDetail tidak ditemukan")
	}
	return productionRequestDetailRes, nil
}

func (repository *ProductionRequestDetailRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB, params map[string][]string) (productionRequestDetailRes []web.ProductionRequestDetailGet, err error) {
	results := db.Table("production_request_details")
	for k, v := range params {
		if v[0] != "" && k != "id" {
			results = results.Where(k+" = ?", v[0])
		}
	}
	results.Find(&productionRequestDetailRes)

	for key, val := range productionRequestDetailRes {
		db.Table("products").Where("id = ?", val.ProductId).First(&productionRequestDetailRes[key].Product.Product)
		db.Table("categories").Where("id = ?", productionRequestDetailRes[key].Product.Product.CategoryId).First(&productionRequestDetailRes[key].Product.Category)
		db.Table("divisions").Where("id = ?", productionRequestDetailRes[key].Product.Category.DivisionId).First(&productionRequestDetailRes[key].Product.Division)
	}
	return productionRequestDetailRes, nil
}

func (repository *ProductionRequestDetailRepositoryImpl) DeleteByProductionRequest(ctx echo.Context, db *gorm.DB, productionRequestId int) (bool, error) {
	productionRequestDetail := domain.ProductionRequestDetail{}
	results := db.Where("production_request_id = ?", productionRequestId).Delete(&productionRequestDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|productionRequestDetail tidak ditemukan")
	}
	return true, nil
}

func (repository *ProductionRequestDetailRepositoryImpl) Pdf(ctx echo.Context, db *gorm.DB, productionRequestId int) (res []web.ProductionRequestDetailReportGet, err error) {
	qry := db.Table("production_request_details")
	qry.Select(`
		production_request_details.current_stock, production_request_details.production_qty,
		products.name product_name, categories.name category_name
	`)
	qry.Joins(`
		JOIN products ON products.id = production_request_details.product_id
		JOIN categories ON categories.id = production_request_details.category_id
	`)
	if productionRequestId != 0 {
		qry.Where("(production_request_details.production_request_id = ?)", productionRequestId)
	}
	qry.Order("production_request_details.id desc")
	qry.Find(&res)
	return res, nil
}