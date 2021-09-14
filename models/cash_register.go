package models
import (
	"time"
)

type CashRegister struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	CashInHand string `json:"cash_in_hand"`
	CreatedBy int `json:"created_by"`
	Status string `json:"status"`
	StoreId string `json:"store_id"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}