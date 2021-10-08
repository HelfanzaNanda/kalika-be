package web

import "kalika-be/models/domain"

type StockOpnameDatatable struct {
	domain.StockOpname
	Store domain.Store `json:"store"`
	Action string `json:"action"`
}