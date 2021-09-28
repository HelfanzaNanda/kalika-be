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
	StoreConsignmentRepository interface{
		Create(ctx echo.Context, db *gorm.DB, storeConsignment *domain.StoreConsignment) (domain.StoreConsignment, error)
		Update(ctx echo.Context, db *gorm.DB, storeConsignment *domain.StoreConsignment) (domain.StoreConsignment, error)
		Delete(ctx echo.Context, db *gorm.DB, storeConsignment *domain.StoreConsignment) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.StoreConsignment, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.StoreConsignment, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.StoreConsignmentDatatable, int64, int64, error)
	}

	StoreConsignmentRepositoryImpl struct {

	}
)

func NewStoreConsignmentRepository() StoreConsignmentRepository {
	return &StoreConsignmentRepositoryImpl{}
}

func (repository StoreConsignmentRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, storeConsignment *domain.StoreConsignment) (domain.StoreConsignment, error) {
	db.Create(&storeConsignment)
	storeConsignmentRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(storeConsignment.Id))
	return storeConsignmentRes, nil
}

func (repository StoreConsignmentRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, storeConsignment *domain.StoreConsignment) (domain.StoreConsignment, error) {
	db.Where("id = ?", storeConsignment.Id).Updates(&storeConsignment)
	storeConsignmentRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(storeConsignment.Id))
	return storeConsignmentRes, nil
}

func (repository StoreConsignmentRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, storeConsignment *domain.StoreConsignment) (bool, error) {
	results := db.Where("id = ?", storeConsignment.Id).Delete(&storeConsignment)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|storeConsignment tidak ditemukan")
	}
	return true, nil
}

func (repository StoreConsignmentRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (storeConsignmentRes domain.StoreConsignment, err error) {
	results := db.Where(key+" = ?", value).First(&storeConsignmentRes)
	if results.RowsAffected < 1 {
		return storeConsignmentRes, errors.New("NOT_FOUND|storeConsignment tidak ditemukan")
	}
	return storeConsignmentRes, nil
}

func (repository StoreConsignmentRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (storeConsignmentRes []domain.StoreConsignment, err error) {
	db.Find(&storeConsignmentRes)
	return storeConsignmentRes, nil
}

func (repository StoreConsignmentRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.StoreConsignmentDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("store_consignments")
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(id = ? OR store_name LIKE ?)", search, "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}

