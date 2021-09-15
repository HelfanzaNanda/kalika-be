package domain

import (
	"time"
)

type UnitConversion struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	UnitId int `json:"unit_id"`
	SmallestUnitId int `json:"smallest_unit_id"`
	Value string `json:"value"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}