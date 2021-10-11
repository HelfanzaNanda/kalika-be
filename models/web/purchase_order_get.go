package web

import "kalika-be/models/domain"

type PurchaseOrderGet struct {
	domain.PurchaseOrder
	SupplierName  string `json:"supplier_name"`
	CreatedByName string `json:"created_by_name"`
}
