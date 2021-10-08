package domain

import (
	"time"
)

type ProductPrice struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	ProductId int `json:"product_id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Type string `json:"type"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}