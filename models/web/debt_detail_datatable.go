package web

import "kalika-be/models/domain"

type DebtDetailDatatable struct {
	domain.DebtDetail
	PaymentMethod string `json:"payment_method"`
	Action string `json:"action"`
}