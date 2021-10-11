package web

import "kalika-be/models/domain"

type ProductGet struct {
	domain.Product
	Division domain.Division `json:"division"`
	Category domain.Category `json:"category"`
}