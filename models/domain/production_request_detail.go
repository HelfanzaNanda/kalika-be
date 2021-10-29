package domain

import (
	"time"
)

type ProductionRequestDetail struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	ProductionRequestId int `json:"production_request_id"`
	ProductId int `json:"product_id"`
	CategoryId int `json:"category_id"`
	CurrentStock float64 `json:"current_stock"`
	ProductionQty float64 `json:"production_qty"`
	Note string `json:"note"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}