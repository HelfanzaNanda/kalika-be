package models
import (
	"time"
)

type SalesConsignment struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	Number string `json:"number"`
	Date time.Time `json:"date"`
	Total float64 `json:"total"`
	StoreConsignmentId int `json:"store_consignment_id"`
	CreatedBy int `json:"created_by"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}