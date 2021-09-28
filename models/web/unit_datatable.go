package web

import "kalika-be/models/domain"

type UnitDatatable struct {
	domain.Unit
	Action string `json:"action"`
}