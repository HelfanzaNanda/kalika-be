package web

import "kalika-be/models/domain"

type CustomOrderGet struct {
	domain.CustomOrder
	PaymentMethodName string `json:"payment_method_name"`
	CreatedByName     string `json:"created_by_name"`
}
