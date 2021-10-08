package web

import (
	"kalika-be/models/domain"
)

type StockOpnameDetailGet struct {
	domain.StockOpnameDetail
	Product domain.Product `json:"product"`
	Store domain.Store `json:"store"`
}