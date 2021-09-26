package web

import "kalika-be/models/domain"

type PaymentMethodDatatable struct {
	domain.PaymentMethod
	Action string `json:"action"`
}