package web

import "kalika-be/models/domain"

type SalesConsignmentGet struct {
	domain.SalesConsignment
	Store domain.Store `json:"store"`
	SalesConsignmentDetails []SalesConsignmentDetailGet `json:"sales_consignment_details"`
	Payment domain.Payment `json:"payment"`
}