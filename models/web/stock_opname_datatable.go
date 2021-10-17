package web

import "kalika-be/models/domain"

type StockOpnameDatatable struct {
	domain.StockOpname
	Store domain.Store `json:"store"`
	CreatedByName string `json:"created_by_name"`
	Action string `json:"action"`
}