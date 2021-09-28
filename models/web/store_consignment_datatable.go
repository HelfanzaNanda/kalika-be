package web

import "kalika-be/models/domain"

type StoreConsignmentDatatable struct {
	domain.StoreConsignment
	Action string `json:"action"`
}