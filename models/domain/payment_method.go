package domain

import (
	"time"
)

type PaymentMethod struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	Name string `json:"name"`
	Deduction float64 `json:"deduction"`
	Description string `json:"description"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}