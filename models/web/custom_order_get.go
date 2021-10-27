package web

import "kalika-be/models/domain"

type CustomOrderGet struct {
	domain.CustomOrder
	PaymentMethodName string             `json:"payment_method_name"`
	CreatedByName     string             `json:"created_by_name"`
	Seller            domain.Seller      `json:"seller"`
	StoreName         string             `json:"store_name"`
	ProductName       string             `json:"product_name"`
	TypeCake          domain.CakeType    `json:"cake_type"`
	VariantCake       domain.CakeVariant `json:"variant_cake"`
}
