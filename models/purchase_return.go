package models
import (
	"time"
)

type PurchaseReturn struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	Model string `json:"model"`
	ModelId int `json:"model_id"`
	Date time.Time `json:"date"`
	Number string `json:"number"`
	CreatedBy int `json:"created_by"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}