package web

import (
	"kalika-be/models/domain"
)

type RecipeDetailGet struct {
	domain.RecipeDetail
	RawMaterial domain.RawMaterial `json:"raw_material"`
}