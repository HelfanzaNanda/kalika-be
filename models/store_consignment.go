package models

import (
	"time"
)

type StoreConsignment struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	StoreName string `json:"store_name"`
	StorePhone string `json:"store_phone"`
	PicName string `json:"pic_name"`
	PicPhone string `json:"pic_phone"`
	Discount float64 `json:"discount"`
	DayOfRules int `json:"day_of_rules"`
	Location string `json:"location"`
	Description string `json:"description"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}