package web

import (
	"kalika-be/models/domain"
)

type SalesConsignmentDetailGet struct {
	domain.SalesConsignmentDetail
	Product domain.Product `json:"product"`
}