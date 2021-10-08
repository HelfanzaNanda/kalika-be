package web

import (
	
)

type DebtPosPost struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	SupplierId int `json:"supplier_id"`
	Model string `json:"model"`
	ModelId int `json:"model_id"`
	Total float64 `json:"total"`
	Debts float64 `json:"debts"`
	Date string `json:"date"`
	Note string `json:"note"`
}