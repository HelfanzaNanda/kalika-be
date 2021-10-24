package web

type ReportSalesConsignmentReportFilterDatatable struct {
	StartDate          string `json:"start_date"`
	EndDate            string `json:"end_date"`
	CreatedBy          int    `json:"created_by"`
	StoreConsignmentId int    `json:"store_consignment_id"`
}
