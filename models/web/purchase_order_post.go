package web

import "kalika-be/models/domain"

type PurchaseOrderPost struct {
	domain.PurchaseOrder
	PurchaseOrderDetails []PurchaseOrderDetailGet `json:"purchase_order_details"`
	Payment domain.Payment `json:"payment"`
}