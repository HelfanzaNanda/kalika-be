package domain

import (
	"time"
)

type PurchaseInvoiceDetail struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	PurchaseOrderId int `json:"purchase_order_id"`
	RawMaterialId int `json:"raw_material_id"`
	Qty int `json:"qty"`
	Price float64 `json:"price"`
	Discount float64 `json:"discount"`
	Total float64 `json:"total"`
	DeliveredQty int `json:"delivered_by"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}