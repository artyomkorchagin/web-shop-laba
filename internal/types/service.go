package types

import "github.com/google/uuid"

type Service struct {
	ServiceID   uuid.UUID `json:"service_id"`
	Name        string    `json:"service_name"`
	Description string    `json:"description"`
	Provider    string    `json:"provider"`
}
