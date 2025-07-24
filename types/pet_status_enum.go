package types

// pet status in the store
type PetStatusEnum string

const (
	PetStatusEnumAvailable PetStatusEnum = "available"
	PetStatusEnumPending   PetStatusEnum = "pending"
	PetStatusEnumSold      PetStatusEnum = "sold"
)
