package web

import "kalika-be/models/domain"

type ReportSalesConsignmentGet struct {
	domain.SalesConsignment
	StoreConsignmentName string `json:"store_consignment_name"`
	CreatedByName        string `json:"created_by_name"`
}
