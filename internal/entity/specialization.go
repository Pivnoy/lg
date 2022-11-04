package entity

import "github.com/google/uuid"

type Specialization struct {
	ID    int64     `json:"id"`
	UUID  uuid.UUID `json:"uuid"`
	Name  string    `json:"name"`
	Value string    `json:"value"`
}
