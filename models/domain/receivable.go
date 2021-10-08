package domain
import (
	"time"
)

type Receivable struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	CustomerId int `json:"customer_id"`
	StoreConsignmentId int `json:"store_consignment_id"`
	Model string `json:"model"`
	ModelId int `json:"model_id"`
	Total float64 `json:"total"`
	Receivables float64 `json:"receivables"`
	Date time.Time `json:"date"`
	Note string `json:"note"`
	CreatedBy int `json:"created_by"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}