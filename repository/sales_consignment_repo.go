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
	SalesConsignmentRepository interface{
		Create(ctx echo.Context, db *gorm.DB, salesConsignment *domain.SalesConsignment) (domain.SalesConsignment, error)
		Update(ctx echo.Context, db *gorm.DB, salesConsignment *domain.SalesConsignment) (domain.SalesConsignment, error)
		Delete(ctx echo.Context, db *gorm.DB, salesConsignment *domain.SalesConsignment) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.SalesConsignment, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.SalesConsignment, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.SalesConsignmentDatatable, int64, int64, error)
	}

	SalesConsignmentRepositoryImpl struct {

	}
)

func NewSalesConsignmentRepository() SalesConsignmentRepository {
	return &SalesConsignmentRepositoryImpl{}
}

func (repository *SalesConsignmentRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, salesConsignment *domain.SalesConsignment) (domain.SalesConsignment, error) {
	db.Create(&salesConsignment)
	salesConsignmentRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(salesConsignment.Id))
	return salesConsignmentRes, nil
}

func (repository *SalesConsignmentRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, salesConsignment *domain.SalesConsignment) (domain.SalesConsignment, error) {
	db.Where("id = ?", salesConsignment.Id).Updates(&salesConsignment)
	salesConsignmentRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(salesConsignment.Id))
	return salesConsignmentRes, nil
}

func (repository *SalesConsignmentRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, salesConsignment *domain.SalesConsignment) (bool, error) {
	results := db.Where("id = ?", salesConsignment.Id).Delete(&salesConsignment)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|salesConsignment tidak ditemukan")
	}
	return true, nil
}

func (repository *SalesConsignmentRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (salesConsignmentRes domain.SalesConsignment, err error) {
	results := db.Where(key+" = ?", value).First(&salesConsignmentRes)
	if results.RowsAffected < 1 {
		return salesConsignmentRes, errors.New("NOT_FOUND|salesConsignment tidak ditemukan")
	}
	return salesConsignmentRes, nil
}

func (repository *SalesConsignmentRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (salesConsignmentRes []domain.SalesConsignment, err error) {
	db.Find(&salesConsignmentRes)
	return salesConsignmentRes, nil
}

func (repository *SalesConsignmentRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (recipeRes []web.SalesConsignmentDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("sales_consignments")
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(sales_consignments.number LIKE ? AND store_consignment.name LIKE ?)", "%"+search+"%", "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	//qry.Joins("JOIN store_consignments ON store_consignments.id = sales_consignments.store_consignment_id")
	qry.Order("sales_consignments.id desc")
	qry.Preload("StoreConsignment").Find(&recipeRes)
	return recipeRes, totalData, totalFiltered, nil
}