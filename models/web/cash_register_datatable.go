package web

import "kalika-be/models/domain"

type CashRegisterDatatable struct {
	domain.CashRegister
	StoreId string `json:"store_id"`
	StoreName string `json:"store_name"`
	Action string `json:"action"`
}