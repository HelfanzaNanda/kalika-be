package web

import "kalika-be/models/domain"

type CategoryGet struct {
	domain.Category
	Division domain.Division `json:"division"`
	TotalProduct int `json:"total_product"`
}