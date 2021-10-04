package web

import "kalika-be/models/domain"

type ExpenseCategoryDatatable struct {
	domain.ExpenseCategory
	Action string `json:"action"`
}