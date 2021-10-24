package web

import "kalika-be/models/domain"

type PurchaseOrderDeliveryPost struct {
	domain.PurchaseOrderDelivery
	PurchaseOrderDeliveryDetails []domain.PurchaseOrderDeliveryDetail `json:"purchase_order_delivery_details"`
}
