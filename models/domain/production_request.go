package domain

import (
	"time"
)

type ProductionRequest struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	Number string `json:"number"`
	StoreId int `json:"store_id"`
	DivisionId int `json:"division_id"`
	Date time.Time `json:"date"`
	Note string `json:"note"`
	CreatedBy int `json:"created_by"`
	Status string `json:"status"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}