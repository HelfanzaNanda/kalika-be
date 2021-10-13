package domain
import (
	"time"
)

type ReceivableDetail struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	ReceivableId int `json:"receivable_id"`
	DatePay time.Time `json:"date_pay"`
	Total float64 `json:"total"`
	PaymentMethodId int `json:"payment_method_id"`
	Note string `json:"note"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}