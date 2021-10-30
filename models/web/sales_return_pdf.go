package web

import "kalika-be/models/domain"

type SalesReturnPdf struct {
	domain.SalesReturn
	CustomerName         string                     `json:"customer_name"`
	StoreConsignmentName string                     `json:"store_consignment_name"`
	CreatedByName        string                     `json:"created_by_name"`
}
