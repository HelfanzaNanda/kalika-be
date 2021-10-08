package web

import "kalika-be/models/domain"

type DebtDatatable struct {
	domain.Debt
	UserName string `json:"user_name"`
	Action string `json:"action"`
}