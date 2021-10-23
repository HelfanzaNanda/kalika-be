package web

import "kalika-be/models/domain"

type ExpenseDatatable struct {
	domain.Expense
	CreatedByName string  `json:"created_by_name"`
	Action        string  `json:"action"`
}
