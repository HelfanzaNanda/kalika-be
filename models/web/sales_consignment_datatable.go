package web

import "kalika-be/models/domain"

type SalesConsignmentDatatable struct {
	domain.SalesConsignment
	StoreConsignment domain.StoreConsignment `json:"store_consignment"`
	CreatedByName string `json:"created_by_name"`
	Action string `json:"action"`
}