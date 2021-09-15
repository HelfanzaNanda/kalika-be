package domain
import (
	"time"
)

type SalesConsignmentDetail struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	SalesConsignmentId int `json:"sales_consignment_id"`
	Qty int `json:"qty"`
	ProductId int `json:"product_id"`
	Total float64 `json:"total"`
	Discount float64 `json:"discount"`
	UnitPrice int `json:"unit_price"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}