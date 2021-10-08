package web

import "kalika-be/models/domain"

type SalesReturnDatatable struct {
	domain.SalesReturn
	Action string `json:"action"`
}