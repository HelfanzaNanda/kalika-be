package web

import "kalika-be/models/domain"

type SalesReturnGet struct {
	domain.SalesReturn
	CustomerName         string                     `json:"customer_name"`
	StoreConsignmentName string                     `json:"store_consignment_name"`
	CreatedByName        string                     `json:"created_by_name"`
	SalesReturnDetail    []domain.SalesReturnDetail `json:"sales_return_details"`
}
