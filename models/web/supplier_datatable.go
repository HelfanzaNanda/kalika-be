package web

import "kalika-be/models/domain"

type SupplierDatatable struct {
	domain.Supplier
	Action string `json:"action"`
}