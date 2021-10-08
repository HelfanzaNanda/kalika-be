package web

import "kalika-be/models/domain"

type ProductLocationGet struct {
	domain.ProductLocation
	Product domain.Product `json:"product"`
	Store domain.Store `json:"store"`
}