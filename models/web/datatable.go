package web

type Datatable struct {
	Draw int `json:"draw"`
	RecordsTotal int64 `json:"recordsTotal"`
	RecordsFiltered int64 `json:"recordsFiltered"`
	Data []interface{} `json:"data"`
	//Data []interface{} `json:"data"`
	Order []interface{} `json:"order"`
}