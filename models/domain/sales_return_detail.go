package domain
import (
	"time"
)

type SalesReturnDetail struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	SalesReturnId int `json:"sales_return_id"`
	ProductId int `json:"product_id"`
	Qty int `json:"qty"`
	Total float64 `json:"total"`
	UnitPrice float64 `json:"unit_price"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}