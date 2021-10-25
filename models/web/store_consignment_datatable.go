package web

import "kalika-be/models/domain"

type StoreConsignmentDatatable struct {
	domain.StoreConsignment
	CreatedByName        string `json:"created_by_name"`
	Action               string `json:"action"`
}
