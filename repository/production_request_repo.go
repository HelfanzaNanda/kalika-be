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
	ProductionRequestRepository interface{
		Create(echo.Context, *gorm.DB, *domain.ProductionRequest) (web.ProductionRequestGet, error)
		Update(echo.Context, *gorm.DB, *domain.ProductionRequest) (web.ProductionRequestGet, error)
		ProcessApprovedProduction(echo.Context, *gorm.DB, *domain.ProductionRequest) (domain.ProductionRequest, error)
		Delete(echo.Context, *gorm.DB, *domain.ProductionRequest) (bool, error)
		FindById(echo.Context, *gorm.DB, string, string) (web.ProductionRequestGet, error)
		FindAll(echo.Context, *gorm.DB) ([]domain.ProductionRequest, error)
		Datatable(echo.Context, *gorm.DB, string, string, string, string) ([]web.ProductionRequestDatatable, int64, int64, error)
		
	}

	ProductionRequestRepositoryImpl struct {

	}
)

func NewProductionRequestRepository() ProductionRequestRepository {
	return &ProductionRequestRepositoryImpl{}
}

func (repository *ProductionRequestRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, productionRequest *domain.ProductionRequest) (web.ProductionRequestGet, error) {
	db.Create(&productionRequest)
	productionRequestRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(productionRequest.Id))
	return productionRequestRes, nil
}

func (repository *ProductionRequestRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, productionRequest *domain.ProductionRequest) (web.ProductionRequestGet, error) {
	db.Where("id = ?", productionRequest.Id).Updates(&productionRequest)
	productionRequestRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(productionRequest.Id))
	return productionRequestRes, nil
}

func (repository *ProductionRequestRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, productionRequest *domain.ProductionRequest) (bool, error) {
	results := db.Where("id = ?", productionRequest.Id).Delete(&productionRequest)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|productionRequest tidak ditemukan")
	}
	return true, nil
}

func (repository *ProductionRequestRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (productionRequestRes web.ProductionRequestGet, err error) {
	results := db.Where(key+" = ?", value).First(&productionRequestRes.ProductionRequest)
	db.Where("id = ?", productionRequestRes.ProductionRequest.StoreId).First(&productionRequestRes.Store)
	db.Where("id = ?", productionRequestRes.ProductionRequest.DivisionId).First(&productionRequestRes.Division)
	db.Model(&domain.User{}).Select("name").Where("id = ?", productionRequestRes.ProductionRequest.CreatedBy).First(&productionRequestRes.CreatedByName)
	if results.RowsAffected < 1 {
		return productionRequestRes, errors.New("NOT_FOUND|productionRequest tidak ditemukan")
	}
	return productionRequestRes, nil
}

func (repository *ProductionRequestRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (productionRequestRes []domain.ProductionRequest, err error) {
	db.Find(&productionRequestRes)
	return productionRequestRes, nil
}

func (repository *ProductionRequestRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (productionRequestRes []web.ProductionRequestDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("production_requests").Select("production_requests.*, users.name as created_by_name")
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(production_requests.number LIKE ?)", "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Joins("JOIN users ON users.id = production_requests.created_by")
	qry.Order("production_requests.id desc")
	qry.Preload("Store").Preload("Division").Find(&productionRequestRes)
	return productionRequestRes, totalData, totalFiltered, nil
}

func (repository *ProductionRequestRepositoryImpl) ProcessApprovedProduction(ctx echo.Context, db *gorm.DB, productionRequest *domain.ProductionRequest) (res domain.ProductionRequest, err error) {
	productionRequestDetails := []domain.ProductionRequestDetail{}
	db.Where("id = ?", productionRequest.Id).Updates(&productionRequest)
	db.Where("production_request_id = ?", productionRequest.Id).Find(&productionRequestDetails)

	for _, val := range productionRequestDetails {
		db.Model(&domain.ProductLocation{}).Where("model = ? AND product_id = ? AND store_id = ?", "Product", val.ProductId, 1).Update("quantity", val.ProductionQty)
	}
	return *productionRequest, nil
}