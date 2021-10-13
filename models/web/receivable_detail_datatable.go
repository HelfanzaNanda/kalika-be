package web

import "kalika-be/models/domain"

type ReceivableDetailDatatable struct {
	domain.ReceivableDetail
	PaymentMethod string `json:"payment_method"`
	Action string `json:"action"`
}