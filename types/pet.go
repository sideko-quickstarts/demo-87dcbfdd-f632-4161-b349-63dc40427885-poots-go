package types

import (
	nullable "pets_go/nullable"
)

// Pet
type Pet struct {
	Category  nullable.Nullable[Category] `json:"category,omitempty"`
	Id        nullable.Nullable[int]      `json:"id,omitempty"`
	Name      string                      `json:"name"`
	PhotoUrls []string                    `json:"photoUrls"`
	// pet status in the store
	Status nullable.Nullable[PetStatusEnum] `json:"status,omitempty"`
	Tags   nullable.Nullable[[]Tag]         `json:"tags,omitempty"`
}
