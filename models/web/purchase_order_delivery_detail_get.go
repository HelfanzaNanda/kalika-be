package web

import (
	"kalika-be/models/domain"
)

type PurchaseOrderDeliveryDetailGet struct {
	domain.PurchaseOrderDeliveryDetail
	RawMaterial domain.RawMaterial `json:"raw_material"`
}