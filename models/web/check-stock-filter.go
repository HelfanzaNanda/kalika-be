package web

type CheckStockFilter struct {
	DivisionId int `json:"division_id"`
	StoreId    int `json:"store_id"`
}
