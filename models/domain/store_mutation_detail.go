package domain

import (
	"time"
)

type StoreMutationDetail struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	StoreMutationId int `json:"store_mutation_id"`
	ProductId int `json:"product_id"`
	UnitPrice float64 `json:"unit_price"`
	Qty float64 `json:"qty"`
	Discount float64 `json:"discount"`
	Note string `json:"note"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}