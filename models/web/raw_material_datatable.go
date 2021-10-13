package web

import "kalika-be/models/domain"

type RawMaterialDatatable struct {
	domain.RawMaterial
	SupplierId int `json:"supplier_id"`
	SupplierName string `json:"supplier_name"`
	UnitId int `json:"unit_id"`
	UnitName string `json:"unit_name"`
	SmallestUnitId int `json:"smallest_unit_id"`
	SmallestUnitName string `json:"smallest_unit_name"`
	StoreId int `json:"store_id"`
	StoreName string `json:"store_name"`
	DivisionName string `json:"division_name"`
	Action string `json:"action"`
}