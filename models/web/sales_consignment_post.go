package web

import "kalika-be/models/domain"

type SalesConsignmentPost struct {
	domain.SalesConsignment
	SalesConsignmentDetails []SalesConsignmentDetailGet `json:"sales_consignment_details"`
	Payment domain.Payment `json:"payment"`
}