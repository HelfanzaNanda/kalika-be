package web

import "kalika-be/models/domain"

type DebtDatatable struct {
	domain.Debt
	SupplierName string `json:"supplier_name"`
	UserName     string `json:"user_name"`
	Action       string `json:"action"`
}
