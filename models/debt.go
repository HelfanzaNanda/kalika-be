package models
import (
	"time"
)

type Debt struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	Model string `json:"model"`
	ModelId int `json:"model_id"`
	Total float64 `json:"total"`
	Debts float64 `json:"debts"`
	Date time.Time `json:"date"`
	Note string `json:"note"`
	CreatedBy int `json:"created_by"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}