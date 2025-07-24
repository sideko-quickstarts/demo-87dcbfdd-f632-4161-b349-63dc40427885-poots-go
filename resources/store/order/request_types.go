package order

import (
	nullable "pets_go/nullable"
	types "pets_go/types"
)

// DeleteRequest
type DeleteRequest struct {
	// ID of the order that needs to be deleted
	OrderId int `json:"orderId"`
}

// GetRequest
type GetRequest struct {
	// ID of order that needs to be fetched
	OrderId int `json:"orderId"`
}

// CreateRequest
type CreateRequest struct {
	Complete nullable.Nullable[bool]   `json:"complete,omitempty"`
	Id       nullable.Nullable[int]    `json:"id,omitempty"`
	PetId    nullable.Nullable[int]    `json:"petId,omitempty"`
	Quantity nullable.Nullable[int]    `json:"quantity,omitempty"`
	ShipDate nullable.Nullable[string] `json:"shipDate,omitempty"`
	// Order Status
	Status nullable.Nullable[types.OrderStatusEnum] `json:"status,omitempty"`
}
