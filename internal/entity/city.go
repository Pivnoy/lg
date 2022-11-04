package entity

import "github.com/google/uuid"

type City struct {
	ID          int64     `json:"id"`
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	CountryUUID uuid.UUID `json:"country_uuid"`
}
