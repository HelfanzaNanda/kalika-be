package web

import "kalika-be/models/domain"

type StoreMutationPost struct {
	domain.StoreMutation
	StoreMutationDetail []domain.StoreMutationDetail `json:"store_mutation_details"`
}