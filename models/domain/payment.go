package domain
import (
	"time"
)

type Payment struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	Number string `json:"number"`
	StoreId int `json:"store_id"`
	Model string `json:"model"`
	ModelId int `json:"model_id"`
	CashRegisterId int `json:"cash_register_id"`
	Total float64 `json:"total"`
	Change int `json:"change"`
	PaymentMethodId int `json:"payment_method_id"`
	PaymentNote string `json:"payment_note"`
	Date time.Time `json:"date"`
	CreatedBy int `json:"created_by"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}