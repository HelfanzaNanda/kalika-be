package web

import "kalika-be/models/domain"

type PurchaseReturnGet struct {
	domain.PurchaseReturn
	CreatedByName string `json:"created_by_name"`
}
