package domain
import (
	"time"
)

type ExpenseDetail struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	ExpenseCategoryId int `json:"expense_category_id"`
	ExpenseId int `json:"expense_id"`
	Amount float64 `json:"amount"`
	Description string `json:"description"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}