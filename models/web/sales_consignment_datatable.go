package web

import "kalika-be/models/domain"

type SalesConsignmentDatatable struct {
	domain.SalesConsignment
	StoreConsignment domain.StoreConsignment `json:"store_consignment"`
	Action string `json:"action"`
}