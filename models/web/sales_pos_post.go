package web

import "kalika-be/models/domain"

type SalesPosPost struct {
	domain.Sale
	SalesDetails []SalesDetailPosGet `json:"sales_details"`
	Payment domain.Payment `json:"payment"`
	Customer domain.Customer `json:"customer"`
}