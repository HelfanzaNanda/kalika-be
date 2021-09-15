package domain

import (
	"time"
)

type PurchaseOrderDeliveryDetail struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	PurchaseOrderId int `json:"purchase_order_id"`
	RawMaterialId int `json:"raw_material_id"`
	DeliveredQty int `json:"delivered_qty"`
	Note string `json:"note"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}