package web

import "kalika-be/models/domain"

type SaleDatatable struct {
	StoreId                int    `json:"store_id"`
	StoreName              string `json:"store_name"`
	CustomerId             int    `json:"customer_id"`
	CustomerName           string `json:"customer_name"`
	CashRegisterId         int    `json:"cash_register_id"`
	CashRegisterCashInHand string `json:"cash_register_cash_in_hand"`
	domain.Sale
	Action string `json:"action"`
}
