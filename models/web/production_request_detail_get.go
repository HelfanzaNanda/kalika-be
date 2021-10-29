package web

import "kalika-be/models/domain"

type ProductionRequestDetailGet struct {
	domain.ProductionRequestDetail
	Product ProductGet `json:"product"`
}