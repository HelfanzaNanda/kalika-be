package web

import "kalika-be/models/domain"

type ProductionRequestPost struct {
	domain.ProductionRequest
	ProductionRequestDetail []domain.ProductionRequestDetail `json:"production_request_details"`
}