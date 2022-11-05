package entity

import "github.com/google/uuid"

type Company struct {
	ID   int64     `json:"id"`
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
	Inn  string    `json:"inn"`
}
