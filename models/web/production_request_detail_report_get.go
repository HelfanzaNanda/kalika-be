package web

type ProductionRequestDetailReportGet struct {
	CurrentStock  int    `json:"current_stock"`
	ProductionQty int    `json:"production_qty"`
	ProductName   string `json:"product_name"`
	CategoryName  string `json:"category_name"`
}
