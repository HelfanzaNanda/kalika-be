package web

import "kalika-be/models/domain"

type CustomOrderDatatable struct {
	StoreId                int    `json:"store_id"`
	StoreName              string `json:"store_name"`
	ProductId                int    `json:"product_id"`
	ProductName              string `json:"product_name"`
	domain.CustomOrder
	Action string `json:"action"`
}
