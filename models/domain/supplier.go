package domain

import (
	"time"
)

type Supplier struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	SalesName string `json:"sales_name"`
	SalesPhone string `json:"sales_phone"`
	Description string `json:"description"`
	TermOfPayment int `json:"term_of_payment"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}