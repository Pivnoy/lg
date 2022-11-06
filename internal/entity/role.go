package entity

import "github.com/google/uuid"

type Role struct {
	ID                 int64     `json:"id"`
	UUID               uuid.UUID `json:"uuid"`
	Name               string    `json:"name"`
	SpecializationUUID uuid.UUID `json:"specialization_uuid"`
}
