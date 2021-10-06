package web

import "kalika-be/models/domain"

type RecipeDatatable struct {
	domain.Recipe
	Product domain.Product `json:"product"`
	Action string `json:"action"`
}