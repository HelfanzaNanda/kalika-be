package web

import "kalika-be/models/domain"

type ProductionRequestGet struct {
	domain.ProductionRequest
	Store domain.Store `json:"store"`
	Division domain.Division `json:"division"`
	CreatedByName string `json:"created_by_name"`
	ProductionRequestDetail []ProductionRequestDetailGet `json:"production_request_details"`
}