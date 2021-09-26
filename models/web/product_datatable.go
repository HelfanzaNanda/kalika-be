package web

import "kalika-be/models/domain"

type ProductDatatable struct {
	domain.Product
	DivisionId int `json:"division_id"`
	DivisionName string `json:"division_name"`
	CategoryId int `json:"category_id"`
	CategoryName string `json:"category_name"`
	CakeTypeId int `json:"cake_type_id"`
	CakeTypeName string `json:"cake_type_name"`
	CakeVariantId int `json:"cake_variant_id"`
	CakeVariantName string `json:"cake_variant_name"`
	Action string `json:"action"`
}