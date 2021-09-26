package web

import "kalika-be/models/domain"

type CategoryDatatable struct {
	DivisionId int `json:"division_id" gorm:"AUTO_INCREMENT"`
	DivisionName string `json:"division_name"`
	domain.Category
	Action string `json:"action"`
}