package types

import (
	nullable "pets_go/nullable"
)

// Tag
type Tag struct {
	Id   nullable.Nullable[int]    `json:"id,omitempty"`
	Name nullable.Nullable[string] `json:"name,omitempty"`
}
