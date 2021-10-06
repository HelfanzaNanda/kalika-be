package web

import "kalika-be/models/domain"

type RecipePost struct {
	domain.Recipe
	RecipeDetail []RecipeDetailGet `json:"recipe_details"`
}