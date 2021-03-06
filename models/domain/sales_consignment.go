package domain
import (
	"time"
)

type SalesConsignment struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	Number string `json:"number"`
	Status string `json:"status"`
	Date time.Time `json:"date"`
	Total float64 `json:"total"`
	StoreConsignmentId int `json:"store_consignment_id"`
	Discount float64 `json:"discount"`
	CreatedBy int `json:"created_by"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}