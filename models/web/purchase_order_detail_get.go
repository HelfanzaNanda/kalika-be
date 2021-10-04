package web

import (
	"kalika-be/models/domain"
)

type PurchaseOrderDetailGet struct {
	domain.PurchaseOrderDetail
	RawMaterial domain.RawMaterial `json:"raw_material"`
}