package web

import "kalika-be/models/domain"

type SalesPosPost struct {
	domain.Sale
	SalesDetails []domain.SalesDetail `json:"sales_details"`
	Payment domain.Payment `json:"payment"`
	Customer domain.Customer `json:"customer"`
}