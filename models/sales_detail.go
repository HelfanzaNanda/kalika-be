package models

import (
	"time"
)

type SalesDetail struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	SalesId int `json:"sales_id"`
	ProductId int `json:"product_id"`
	Qty int `json:"qty"`
	DiscountPercentage int `json:"discount_percentage"`
	DiscountValue int `json:"discount_value"`
	Total float64 `json:"total"`
	UnitPrice int `json:"unit_price"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}