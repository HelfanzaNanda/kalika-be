package domain

import (
	"time"
)

type PurchaseInvoice struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	Number string `json:"number"`
	SupplierId int `json:"supplier_id"`
	PurchaseOrderId int `json:"purchase_order_id"`
	Date time.Time `json:"date"`
	Status string `json:"status"`
	Discount float64 `json:"discount"`
	Total float64 `json:"total"`
	CreatedBy int `json:"created_by"`
	ApprovedBy int `json:"approved_by"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}