package web

import "kalika-be/models/domain"

type StoreMutationDetailGet struct {
	domain.StoreMutationDetail
	Product ProductGet `json:"product"`
}