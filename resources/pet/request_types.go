package pet

import (
	os "os"
	nullable "pets_go/nullable"
	types "pets_go/types"
)

// DeleteRequest
type DeleteRequest struct {
	// Pet id to delete
	PetId int `json:"petId"`
}

// FindByStatusRequest
type FindByStatusRequest struct {
	// Status values that need to be considered for filter
	Status nullable.Nullable[types.PetFindByStatusStatusEnum] `json:"status,omitempty"`
}

// GetRequest
type GetRequest struct {
	// ID of pet to return
	PetId int `json:"petId"`
}

// CreateRequest
type CreateRequest struct {
	Category  nullable.Nullable[types.Category] `json:"category,omitempty"`
	Id        nullable.Nullable[int]            `json:"id,omitempty"`
	Name      string                            `json:"name"`
	PhotoUrls []string                          `json:"photoUrls"`
	// pet status in the store
	Status nullable.Nullable[types.PetStatusEnum] `json:"status,omitempty"`
	Tags   nullable.Nullable[[]types.Tag]         `json:"tags,omitempty"`
}

// UploadImageRequest
type UploadImageRequest struct {
	Data os.File `json:"data"`
	// ID of pet to update
	PetId int `json:"petId"`
	// Additional Metadata
	AdditionalMetadata nullable.Nullable[string] `json:"additionalMetadata,omitempty"`
}

// UpdateRequest
type UpdateRequest struct {
	Category  nullable.Nullable[types.Category] `json:"category,omitempty"`
	Id        nullable.Nullable[int]            `json:"id,omitempty"`
	Name      string                            `json:"name"`
	PhotoUrls []string                          `json:"photoUrls"`
	// pet status in the store
	Status nullable.Nullable[types.PetStatusEnum] `json:"status,omitempty"`
	Tags   nullable.Nullable[[]types.Tag]         `json:"tags,omitempty"`
}
