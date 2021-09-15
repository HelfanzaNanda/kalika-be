package domain

import (
	"time"
)

type CustomOrder struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	StoreId int `json:"store_id"`
	ProductId int `json:"product_id"`
	CakeCharacter string `json:"cake_character"`
	CakeShape string `json:"cake_shape"`
	CakeSize int `json:"cake_size"`
	DeliveryDate time.Time `json:"delivery_date"`
	Price float64 `json:"price"`
	Other string `json:"other"`
	CakeCustomName string `json:"cake_custom_name"`
	AdditionalPrice float64 `json:"additional_price"`
	Candle string `json:"candle"`
	Discount float64 `json:"discount"`
	DeliveryCost float64 `json:"delivery_cost"`
	RecipientName string `json:"recipient_name"`
	RecipientAddress string `json:"recipient_address"`
	RecipientMethod string `json:"recipient_method"`
	RecipientStatus string `json:"recipient_status"`
	ShipmentType string `json:"shipment_type"`
	Note string `json:"note"`
	Status string `json:"status"`
	CreatedBy int `json:"created_by"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}