package web

type ReportStockOpnameGet struct {
	BookStock     int    `json:"book_stock"`
	PhysicalStock int    `json:"physical_stock"`
	ProductName   string `json:"product_name"`
	CategoryName  string `json:"category_name"`
	MinimumStock  int    `json:"minimum_stock"`
}
