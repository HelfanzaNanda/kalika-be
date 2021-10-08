package web

import "kalika-be/models/domain"

type StockOpnamePost struct {
	domain.StockOpname
	StockOpnameDetail []domain.StockOpnameDetail `json:"stock_opname_details"`
}