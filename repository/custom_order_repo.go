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
	CustomOrderRepository interface{
		Create(ctx echo.Context, db *gorm.DB, customOrder *domain.CustomOrder) (domain.CustomOrder, error)
		Update(ctx echo.Context, db *gorm.DB, customOrder *domain.CustomOrder) (domain.CustomOrder, error)
		Delete(ctx echo.Context, db *gorm.DB, customOrder *domain.CustomOrder) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.CustomOrder, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.CustomOrder, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.CustomOrderDatatable, int64, int64, error)
	}

	CustomOrderRepositoryImpl struct {

	}
)

func NewCustomOrderRepository() CustomOrderRepository {
	return &CustomOrderRepositoryImpl{}
}

func (repository CustomOrderRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, customOrder *domain.CustomOrder) (domain.CustomOrder, error) {
	db.Create(&customOrder)
	customOrderRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(customOrder.Id))
	return customOrderRes, nil
}

func (repository CustomOrderRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, customOrder *domain.CustomOrder) (domain.CustomOrder, error) {
	db.Where("id = ?", customOrder.Id).Updates(&customOrder)
	customOrderRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(customOrder.Id))
	return customOrderRes, nil
}

func (repository CustomOrderRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, customOrder *domain.CustomOrder) (bool, error) {
	results := db.Where("id = ?", customOrder.Id).Delete(&customOrder)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|customOrder tidak ditemukan")
	}
	return true, nil
}

func (repository CustomOrderRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (customOrderRes domain.CustomOrder, err error) {
	results := db.Where(key+" = ?", value).First(&customOrderRes)
	if results.RowsAffected < 1 {
		return customOrderRes, errors.New("NOT_FOUND|customOrder tidak ditemukan")
	}
	return customOrderRes, nil
}

func (repository CustomOrderRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (customOrderRes []domain.CustomOrder, err error) {
	db.Find(&customOrderRes)
	return customOrderRes, nil
}

func (repository CustomOrderRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.CustomOrderDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("custom_orders").
		Select(`
			custom_orders.*,
			products.id product_id, products.name product_name, 
			stores.id store_id, stores.name store_name,
			users.name created_by_name
		`).
		Joins(`
			left join stores on stores.id = custom_orders.store_id
			left join products on products.id = custom_orders.product_id
			left join users on users.id = custom_orders.created_by
		`)

	qry.Count(&totalData)
	if search != "" {
		qry.Where("(custom_orders.id = ? OR custom_orders.cake_character LIKE ?)", search, "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("custom_orders.id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}