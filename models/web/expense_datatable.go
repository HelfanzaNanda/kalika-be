package web

import "kalika-be/models/domain"

type ExpenseDatatable struct {
	domain.Expense
	Action string `json:"action"`
}