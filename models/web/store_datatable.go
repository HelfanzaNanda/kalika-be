package web

import "kalika-be/models/domain"

type StoreDatatable struct {
	domain.Store
	Action string `json:"action"`
}