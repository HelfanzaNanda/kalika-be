package web

import "kalika-be/models/domain"

type ReceivableDatatable struct {
	domain.Receivable
	CustomerName         string `json:"customer_name"`
	StoreConsignmentName string `json:"store_consignment_name"`
	UserName             string `json:"user_name"`
	Action               string `json:"action"`
}
