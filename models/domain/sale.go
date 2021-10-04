package domain

import (
	"time"
)

type Sale struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	Number string `json:"number"`
	StoreId int `json:"store_id"`
	CashRegisterId int `json:"cash_register_id"`
	CustomerId int `json:"customer_id"`
	DiscountPercentage int `json:"discount_percentage"`
	DiscountValue int `json:"discount_value"`
	Total float64 `json:"total"`
	PaymentStatus string `json:"payment_status"`
	SaleStatus string `json:"sale_status"`
	Note string `json:"note"`
	CreatedBy int `json:"created_by"`
	CustomerPay float64 `json:"customer_pay"`
	CustomerChange float64 `json:"customer_change"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}