package web

import (
	"time"
)

type ReportLedgerReceivable struct {
	Date time.Time `json:"date"`
	Number string `json:"number"`
	Customer string `json:"customer"`
	Description string `json:"description"`
	Debit float64 `json:"debit"`
	Credit float64 `json:"credit"`
	Balance float64 `json:"balance"`
}
