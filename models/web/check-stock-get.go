package web

type CheckStockGet struct {
	ProductName  string `json:"product_name"`
	DivisionName string `json:"division_name"`
	CategoryName string `json:"category_name"`
	Qty          int    `json:"qty"`
	MinimumStock int    `json:"minimum_stock"`
}
