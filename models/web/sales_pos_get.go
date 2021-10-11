package web

import "kalika-be/models/domain"

type SalesPosGet struct {
	domain.Sale
	StoreName     string `json:"store_name"`
	CustomerName  string `json:"customer_name"`
	CashInHand    string `json:"cash_in_hand"`
	CreatedByName string `json:"created_by_name"`
}
