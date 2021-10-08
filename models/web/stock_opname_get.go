package web

import "kalika-be/models/domain"

type StockOpnameGet struct {
	domain.StockOpname
	Store domain.Store `json:"store"`
	StockOpnameDetail []StockOpnameDetailGet `json:"stock_opname_details"`
}