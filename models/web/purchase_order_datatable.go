package web

import "kalika-be/models/domain"

type PurchaseOrderDatatable struct {
	domain.PurchaseOrder
	SupplierName  string `json:"supplier_name"`
	CreatedByName string `json:"created_by_name"`
	Action        string `json:"action"`
}
