package web

type ProfitLossReport struct {
	Sales float64 `json:"sales"`
	TotalCogs float64 `json:"total_cogs"`
	CustomOrder float64 `json:"custom_order"`
	SalesConsignment float64 `json:"sales_consignment"`
	ReceivablePayment float64 `json:"receivable_payment"`
	DebtPayment float64 `json:"debt_payment"`
	SalesReturn float64 `json:"sales_return"`
	TotalCost float64 `json:"total_cost"`
	Cost []ProfitLossCost `json:"costs"`
}

type ProfitLossCost struct {
	Name string `json:"name"`
	Total float64 `json:"total"`
}
