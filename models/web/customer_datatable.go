package web

import "kalika-be/models/domain"

type CustomerDatatable struct {
	domain.Customer
	Action string `json:"action"`
}