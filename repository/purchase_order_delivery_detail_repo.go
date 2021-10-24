package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	PurchaseOrderDeliveryDetailRepository interface{
		Create(ctx echo.Context, db *gorm.DB, purchaseOrderDeliveryDetail []domain.PurchaseOrderDeliveryDetail) ([]domain.PurchaseOrderDeliveryDetail, error)
		Update(ctx echo.Context, db *gorm.DB, purchaseOrderDeliveryDetail *domain.PurchaseOrderDeliveryDetail) (domain.PurchaseOrderDeliveryDetail, error)
		Delete(ctx echo.Context, db *gorm.DB, purchaseOrderDeliveryDetail *domain.PurchaseOrderDeliveryDetail) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.PurchaseOrderDeliveryDetail, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.PurchaseOrderDeliveryDetail, error)
	}

	PurchaseOrderDeliveryDetailRepositoryImpl struct {

	}
)

func NewPurchaseOrderDeliveryDetailRepository() PurchaseOrderDeliveryDetailRepository {
	return &PurchaseOrderDeliveryDetailRepositoryImpl{}
}

func (repository PurchaseOrderDeliveryDetailRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, purchaseOrderDeliveryDetail []domain.PurchaseOrderDeliveryDetail) (res []domain.PurchaseOrderDeliveryDetail, err error) {
	purchaseOrderId := helpers.StringToInt(ctx.Get("purchase_order_id").(string))
	purchaseOrderDetails := []domain.PurchaseOrderDetail{}
	for _, val := range purchaseOrderDeliveryDetail {
		val.PurchaseOrderDeliveryId = helpers.StringToInt(ctx.Get("purchase_order_delivery_id").(string))
		db.Table("purchase_order_delivery_details").Select("purchase_order_delivery_id", "raw_material_id", "delivered_qty", "note").Create(&val)
		db.Model(&domain.PurchaseOrderDetail{}).Where("purchase_order_id = ? AND raw_material_id = ?", purchaseOrderId, val.RawMaterialId).Update("delivered_qty", gorm.Expr("delivered_qty + ?", val.DeliveredQty))

		res = append(res, val)
	}

	countDeliveredQty := 0
	db.Model(&purchaseOrderDetails).Where("purchase_order_id", purchaseOrderId).Find(&purchaseOrderDetails)
	for _, val := range purchaseOrderDetails {
		if val.Qty > val.DeliveredQty {
			db.Model(&domain.PurchaseOrder{}).Where("id", purchaseOrderId).Update("status", "partially_delivered")
			break
		} else {
			countDeliveredQty++
		}
	}

	if len(purchaseOrderDetails) == countDeliveredQty {
		db.Model(&domain.PurchaseOrder{}).Where("id", purchaseOrderId).Update("status", "completed")
	}

	return res, nil
}

func (repository PurchaseOrderDeliveryDetailRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, purchaseOrderDeliveryDetail *domain.PurchaseOrderDeliveryDetail) (domain.PurchaseOrderDeliveryDetail, error) {
	db.Where("id = ?", purchaseOrderDeliveryDetail.Id).Updates(&purchaseOrderDeliveryDetail)
	purchaseOrderDeliveryDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(purchaseOrderDeliveryDetail.Id))
	return purchaseOrderDeliveryDetailRes, nil
}

func (repository PurchaseOrderDeliveryDetailRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, purchaseOrderDeliveryDetail *domain.PurchaseOrderDeliveryDetail) (bool, error) {
	results := db.Where("id = ?", purchaseOrderDeliveryDetail.Id).Delete(&purchaseOrderDeliveryDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|purchaseOrderDeliveryDetail tidak ditemukan")
	}
	return true, nil
}

func (repository PurchaseOrderDeliveryDetailRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (purchaseOrderDeliveryDetailRes domain.PurchaseOrderDeliveryDetail, err error) {
	results := db.Where(key+" = ?", value).First(&purchaseOrderDeliveryDetailRes)
	if results.RowsAffected < 1 {
		return purchaseOrderDeliveryDetailRes, errors.New("NOT_FOUND|purchaseOrderDeliveryDetail tidak ditemukan")
	}
	return purchaseOrderDeliveryDetailRes, nil
}

func (repository PurchaseOrderDeliveryDetailRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (purchaseOrderDeliveryDetailRes []domain.PurchaseOrderDeliveryDetail, err error) {
	db.Find(&purchaseOrderDeliveryDetailRes)
	return purchaseOrderDeliveryDetailRes, nil
}

