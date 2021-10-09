package web

import "kalika-be/models/domain"

type PurchaseReturnDatatable struct {
	domain.PurchaseReturn
	CreatedByName string `json:"created_by_name"`
	Action        string `json:"action"`
}
