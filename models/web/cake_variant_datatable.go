package web

import "kalika-be/models/domain"

type CakeVariantDatatable struct {
	domain.CakeVariant
	Action string `json:"action"`
}