package web

import "kalika-be/models/domain"

type SellerDatatable struct {
	domain.Seller
	Action string `json:"action"`
}