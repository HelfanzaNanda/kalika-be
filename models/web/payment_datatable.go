package web

import "kalika-be/models/domain"

type PaymentDatatable struct {
	domain.Payment
	PaymentMethod string `json:"payment_method"`
	CreatedByName string `json:"created_by_name"`
	StoreName string `json:"store_name"`
	Action string `json:"action"`
}