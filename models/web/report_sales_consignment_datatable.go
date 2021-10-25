package web

import "kalika-be/models/domain"

type ReportSalesConsignmentDatatable struct {
	domain.SalesConsignment
	StoreConsignmentName string `json:"store_consignment_name"`
	CreatedByName        string `json:"created_by_name"`
	Action               string `json:"action"`
}
