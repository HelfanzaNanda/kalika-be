package web

import "kalika-be/models/domain"

type ProductLocationGet struct {
	domain.ProductLocation
	Product ProductGet `json:"product"`
	Store domain.Store `json:"store"`
}