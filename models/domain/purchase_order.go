package domain

import (
	"time"
)

type PurchaseOrder struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	Number string `json:"number"`
	SupplierId int `json:"supplier_id"`
	Date time.Time `json:"date"`
	Status string `json:"status"`
	Discount float64 `json:"discount"`
	Total float64 `json:"total"`
	CreatedBy int `json:"created_by"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}