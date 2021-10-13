package web

import "kalika-be/models/domain"

type PurchaseReturnPost struct {
	domain.PurchaseReturn
	Date                  string                        `json:"date"`
	PurchaseReturnDetails []domain.PurchaseReturnDetail `json:"purchase_return_details"`
}
