package domain

import (
	"time"
)

type Recipe struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	ProductId int `json:"product_id"`
	Qty int `json:"qty"`
	Total float64 `json:"total"`
	TotalCogs float64 `json:"total_cogs"`
	OverheadPercentage float64 `json:"overhead_percentage"`
	OverheadPrice float64 `json:"overhead_price"`
	RecommendationPercentage float64 `json:"recommendation_percentage"`
	RecommendationPrice float64 `json:"recommendation_price"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}