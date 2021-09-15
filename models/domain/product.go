package domain

import "time"

type Product struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	Name string `json:"name"`
	StockMinimum int `json:"stock_minimum"`
	ProductionMinimum int `json:"production_minimum"`
	DivisionId int `json:"division_id"`
	CategoryId int `json:"category_id"`
	Active bool `json:"active"`
	IsCustomPrice bool `json:"is_custom_price"`
	IsCustomProduct bool `json:"is_custom_product"`
	CakeTypeId int `json:"cake_type_id"`
	CakeVariantId int `json:"cake_variant_id"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}