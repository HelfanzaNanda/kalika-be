package web

type CustomOrderReportFilterDatatable struct {
	//StoreId         int    `json:"store_id"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	CreatedBy       int    `json:"created_by"`
	PaymentMethodId int    `json:"payment_method_id"`
}
