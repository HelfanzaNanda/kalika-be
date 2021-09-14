package models

import (
	"time"
)

type PurchaseOrderDelivery struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	PurchaseOrderId int `json:"purchase_order_id"`
	Number string `json:"number"`
	Date time.Time `json:"date"`
	InvoiceNumber string `json:"invoice_number"`
	CreatedBy int `json:"created_by"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}