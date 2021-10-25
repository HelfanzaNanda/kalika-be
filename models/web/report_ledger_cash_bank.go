package web

import (
	"time"
)

type ReportLedgerCashBank struct {
	Date time.Time `json:"date"`
	Number string `json:"number"`
	PaymentMethod string `json:"payment_method"`
	Type string `json:"type"`
	Description string `json:"description"`
	Debit float64 `json:"debit"`
	Credit float64 `json:"credit"`
	Balance float64 `json:"balance"`
}
