package web

import "kalika-be/models/domain"

type ProductPosPost struct {
	domain.Product
	ProductPrices []domain.ProductPrice `json:"product_prices"`
	ProductLocations []domain.ProductLocation `json:"product_locations"`
}