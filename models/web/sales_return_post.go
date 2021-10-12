package web

import "kalika-be/models/domain"

type SalesReturnPost struct {
	domain.SalesReturn
	SalesReturnDetails []domain.SalesReturnDetail `json:"sales_return_details"`
}
