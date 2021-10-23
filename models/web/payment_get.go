package web

import "kalika-be/models/domain"

type PaymentGet struct {
	domain.Payment
	PaymentMethodName string `json:"payment_method_name"`
}
