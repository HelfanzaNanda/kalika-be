package domain

import (
	"time"
)

type RecipeDetail struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	RecipeId int `json:"recipe_id"`
	RawMaterialId int `json:"raw_material_id"`
	Price float64 `json:"price"`
	UnitId int `json:"unit_id"`
	Quantity float64 `json:"quantity"`
	Total float64 `json:"total"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}