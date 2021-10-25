package web

import (
	"time"
)

type ReportLedgerDebt struct {
	Date time.Time `json:"date"`
	Number string `json:"number"`
	Supplier string `json:"supplier"`
	Description string `json:"description"`
	Debit float64 `json:"debit"`
	Credit float64 `json:"credit"`
	Balance float64 `json:"balance"`
}
