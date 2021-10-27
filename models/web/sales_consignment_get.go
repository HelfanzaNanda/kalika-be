package web

import "kalika-be/models/domain"

type SalesConsignmentGet struct {
	domain.SalesConsignment
	StoreConsignment domain.StoreConsignment `json:"store_consignment"`
	SalesConsignmentDetails []SalesConsignmentDetailGet `json:"sales_consignment_details"`
	Payment domain.Payment `json:"payment"`
	CreateByName string `json:"create_by_name"`
}