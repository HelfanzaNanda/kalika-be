package web

import "kalika-be/models/domain"

type CakeTypeDatatable struct {
	domain.CakeType
	Action string `json:"action"`
}