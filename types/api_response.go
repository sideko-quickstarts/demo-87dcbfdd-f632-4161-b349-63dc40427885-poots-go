package types

import (
	nullable "pets_go/nullable"
)

// ApiResponse
type ApiResponse struct {
	Code    nullable.Nullable[int]    `json:"code,omitempty"`
	Message nullable.Nullable[string] `json:"message,omitempty"`
	Type    nullable.Nullable[string] `json:"type,omitempty"`
}
