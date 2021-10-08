package web

import "kalika-be/models/domain"

type ProductPosPost struct {
	domain.Product
	ProductPrices []domain.ProductPrice `json:"product_prices"`
}