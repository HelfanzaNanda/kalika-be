package web

import "kalika-be/models/domain"

type PurchaseOrderDatatable struct {
	domain.PurchaseOrder
	Action string `json:"action"`
}