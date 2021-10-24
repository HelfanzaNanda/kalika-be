package web

import "kalika-be/models/domain"

type PurchaseOrderDeliveryGet struct {
	domain.PurchaseOrderDelivery
	PurchaseOrder  string `json:"purchase_order"`
	CreatedByName string `json:"created_by_name"`
	PurchaseOrderDeliveryDetail []PurchaseOrderDeliveryDetailGet `json:"purchase_order_delivery"`
}
