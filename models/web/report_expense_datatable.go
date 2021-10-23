package web

type ReportExpenseDatatable struct {
	Number       string  `json:"number"`
	CategoryName string  `json:"category_name"`
	Total        float64 `json:"total"`
}
