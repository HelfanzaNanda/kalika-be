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
	PurchaseOrderDetailRepository interface{
		Create(ctx echo.Context, db *gorm.DB, purchaseOrder *web.PurchaseOrderPost) ([]web.PurchaseOrderDetailGet, error)
		Update(ctx echo.Context, db *gorm.DB, purchaseOrderDetail *domain.PurchaseOrderDetail) (domain.PurchaseOrderDetail, error)
		Delete(ctx echo.Context, db *gorm.DB, purchaseOrderDetail *domain.PurchaseOrderDetail) (bool, error)
		DeleteByPurchaseOrder(ctx echo.Context, db *gorm.DB, purchaseOrderId int) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.PurchaseOrderDetail, error)
		FindAll(ctx echo.Context, db *gorm.DB, params map[string][]string) ([]web.PurchaseOrderDetailGet, error)
	}

	PurchaseOrderDetailRepositoryImpl struct {

	}
)

func NewPurchaseOrderDetailRepository() PurchaseOrderDetailRepository {
	return &PurchaseOrderDetailRepositoryImpl{}
}

func (repository PurchaseOrderDetailRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, purchaseOrder *web.PurchaseOrderPost) (res []web.PurchaseOrderDetailGet, err error) {
	for _, val := range purchaseOrder.PurchaseOrderDetails {
		val.PurchaseOrderId = purchaseOrder.Id
		val.Total = float64(val.Qty) * val.PurchaseOrderDetail.Price
		val.DeliveredQty = 0
		db.Table("purchase_order_details").Select("purchase_order_id", "raw_material_id", "qty", "price", "discount", "total", "delivered_qty").Create(&val)
		res = append(res, val)
	}

	return res, nil
}

func (repository PurchaseOrderDetailRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, purchaseOrderDetail *domain.PurchaseOrderDetail) (domain.PurchaseOrderDetail, error) {
	db.Where("id = ?", purchaseOrderDetail.Id).Updates(&purchaseOrderDetail)
	purchaseOrderDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(purchaseOrderDetail.Id))
	return purchaseOrderDetailRes, nil
}

func (repository PurchaseOrderDetailRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, purchaseOrderDetail *domain.PurchaseOrderDetail) (bool, error) {
	results := db.Where("id = ?", purchaseOrderDetail.Id).Delete(&purchaseOrderDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|purchaseOrderDetail tidak ditemukan")
	}
	return true, nil
}

func (repository PurchaseOrderDetailRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (purchaseOrderDetailRes domain.PurchaseOrderDetail, err error) {
	results := db.Where(key+" = ?", value).First(&purchaseOrderDetailRes)
	if results.RowsAffected < 1 {
		return purchaseOrderDetailRes, errors.New("NOT_FOUND|purchaseOrderDetail tidak ditemukan")
	}
	return purchaseOrderDetailRes, nil
}

func (repository PurchaseOrderDetailRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB, params map[string][]string) (purchaseOrderDetailRes []web.PurchaseOrderDetailGet, err error) {
	results := db.Table("purchase_order_details").Preload("RawMaterial")
	for k, v := range params {
		if v[0] != "" && k != "id" {
			results = results.Where(k+" = ?", v[0])
		}
	}

	results.Find(&purchaseOrderDetailRes)
	return purchaseOrderDetailRes, nil
}

func (repository PurchaseOrderDetailRepositoryImpl) DeleteByPurchaseOrder(ctx echo.Context, db *gorm.DB, purchaseOrderId int) (bool, error) {
	purchaseOrderDetail := domain.PurchaseOrderDetail{}
	results := db.Where("purchase_order_id = ?", purchaseOrderId).Delete(&purchaseOrderDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|salesDetail tidak ditemukan")
	}
	return true, nil
}