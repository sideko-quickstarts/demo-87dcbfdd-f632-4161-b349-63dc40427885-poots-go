package types

import (
	nullable "pets_go/nullable"
)

// Order
type Order struct {
	Complete nullable.Nullable[bool]   `json:"complete,omitempty"`
	Id       nullable.Nullable[int]    `json:"id,omitempty"`
	PetId    nullable.Nullable[int]    `json:"petId,omitempty"`
	Quantity nullable.Nullable[int]    `json:"quantity,omitempty"`
	ShipDate nullable.Nullable[string] `json:"shipDate,omitempty"`
	// Order Status
	Status nullable.Nullable[OrderStatusEnum] `json:"status,omitempty"`
}
