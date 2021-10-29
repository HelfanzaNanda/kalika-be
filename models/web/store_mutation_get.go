package web

import "kalika-be/models/domain"

type StoreMutationGet struct {
	domain.StoreMutation
	OriginStore domain.Store `json:"origin_store"`
	DestinationStore domain.Store `json:"destination_store_store"`
	CreatedByName string `json:"created_by_name"`
	StoreMutationDetail []StoreMutationDetailGet `json:"store_mutation_details"`
}