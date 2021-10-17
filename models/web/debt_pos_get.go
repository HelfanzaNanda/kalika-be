package web

import "kalika-be/models/domain"

type DebtPosGet struct {
	domain.Debt
	CreatedByName string `json:"created_by_name"`
	SupplierName  string `json:"supplier_name"`
}
