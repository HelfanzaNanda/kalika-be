package web

import "kalika-be/models/domain"

type PurchaseReturnDatatable struct {
	domain.PurchaseReturn
	Action string `json:"action"`
}