package web

import "kalika-be/models/domain"

type ReceivableDatatable struct {
	domain.Receivable
	UserName string `json:"user_name"`
	Action string `json:"action"`
}