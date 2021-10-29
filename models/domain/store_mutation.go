package domain

import (
	"time"
)

type StoreMutation struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	Number string `json:"number"`
	OriginStoreId int `json:"origin_store_id"`
	DestinationStoreId int `json:"destination_store_id"`
	Note string `json:"note"`
	CreatedBy int `json:"created_by"`
	Status string `json:"status"`
	Type string `json:"type"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}