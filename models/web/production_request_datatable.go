package web

import "kalika-be/models/domain"

type ProductionRequestDatatable struct {
	domain.ProductionRequest
	Store domain.Store `json:"store"`
	Division domain.Division `json:"division"`
	CreatedByName string `json:"created_by_name"`
	Action string `json:"action"`
}