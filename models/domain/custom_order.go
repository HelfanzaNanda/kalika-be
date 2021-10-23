package domain

import (
	"time"
)

type CustomOrder struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	Number string `json:"number"`
	StoreId int `json:"store_id"`
	ProductId int `json:"product_id"`
	SellerId int `json:"seller_id"`
	TypeCakeId int `json:"type_cake_id"`
	CakeCharacter string `json:"cake_character"`
	CakeShape string `json:"cake_shape"`
	CakeSize string `json:"cake_size"`
	VariantCakeId int `json:"variant_cake_id"`
	DeliveryDate time.Time `json:"delivery_date"`
	Other string `json:"other"`
	CakeCustomName string `json:"cake_custom_name"`
	Candle string `json:"candle"`
	ShipmentType string `json:"shipment_type"`
	Status string `json:"status"`
	ProductionNote string `json:"production_note"`
	CustomerName string `json:"customer_name"`
	CustomerPhone string `json:"customer_phone"`
	CustomerAddress string `json:"customer_address"`
	CustomerAddressDescription string `json:"customer_address_description"`
	RecipientName string `json:"recipient_name"`
	RecipientPhone string `json:"recipient_phone"`
	RecipientAddress string `json:"recipient_address"`
	RecipientAddressDetail string `json:"recipient_address_detail"`
	Price float64 `json:"price"`
	AdditionalPrice float64 `json:"additional_price"`
	Discount float64 `json:"discount"`
	DeliveryCost float64 `json:"delivery_cost"`
	Total float64 `json:"total"`
	PaymentMethodId int `json:"payment_method_id"`
	DownPayment float64 `json:"down_payment"`
	PaymentNote string `json:"payment_note"`
	CreatedBy int `json:"created_by"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}