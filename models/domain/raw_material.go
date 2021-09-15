package domain

import (
	"time"
)

type RawMaterial struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	Name string `json:"name"`
	SupplierId int `json:"supplier_id"`
	Price float64 `json:"price"`
	UnitId int `json:"unit_id"`
	SmallestUnitId int `json:"smallest_unit_id"`
	Stock int `json:"stock"`
	StoreId int `json:"store_id"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}