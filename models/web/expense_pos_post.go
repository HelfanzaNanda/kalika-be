package web

import "kalika-be/models/domain"

type ExpensePosPost struct {
	domain.Expense
	ExpenseDetails []domain.ExpenseDetail `json:"expense_details"`
}