package domain

import (
	"time"
)

type ProductLocation struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	Model string `json:"model"`
	ProductId int `json:"product_id"`
	StoreId int `json:"store_id"`
	Quantity float64 `json:"quantity"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}