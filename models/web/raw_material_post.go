package web

import "kalika-be/models/domain"

type RawMaterialPost struct {
	domain.RawMaterial
	ProductLocations []domain.ProductLocation `json:"product_locations"`
}