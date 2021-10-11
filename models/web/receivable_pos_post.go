package web

import (
	
)

type ReceivablePosPost struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	CustomerId int `json:"customer_id"`
	StoreConsignmentId int `json:"store_consignment_id"`
	Model string `json:"model"`
	ModelId int `json:"model_id"`
	Total float64 `json:"total"`
	Receivables float64 `json:"receivables"`
	Date string `json:"date"`
	Note string `json:"note"`
}