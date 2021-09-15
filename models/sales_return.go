package models
import (
	"time"
)

type SalesReturn struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	Number string `json:"number"`
	Model string `json:"model"`
	ModelId int `json:"model_id"`
	Type string `json:"type"`
	CreatedBy int `json:"created_by"`
	CustomerId int `json:"customer_id"`
	StoreConsignmentId int `json:"store_consignment_id"`
	Total float64 `json:"total"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}