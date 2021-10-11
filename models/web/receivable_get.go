package web

import "time"

type ReceivableGet struct {
	Id                   int     `json:"id" gorm:"AUTO_INCREMENT"`
	CustomerName         string  `json:"customer_name"`
	StoreConsignmentName string  `json:"store_consignment_name"`
	CreatedByName        string  `json:"created_by_name"`
	Model                string  `json:"model"`
	ModelId              int     `json:"model_id"`
	Total                float64 `json:"total"`
	Receivables          float64 `json:"receivables"`
	Date                 time.Time  `json:"date"`
	Note                 string  `json:"note"`
}
