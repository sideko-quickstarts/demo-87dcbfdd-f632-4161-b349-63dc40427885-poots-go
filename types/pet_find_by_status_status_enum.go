package types

// Status values that need to be considered for filter
type PetFindByStatusStatusEnum string

const (
	PetFindByStatusStatusEnumAvailable PetFindByStatusStatusEnum = "available"
	PetFindByStatusStatusEnumPending   PetFindByStatusStatusEnum = "pending"
	PetFindByStatusStatusEnumSold      PetFindByStatusStatusEnum = "sold"
)
