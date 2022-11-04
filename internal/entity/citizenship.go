package entity

import "github.com/google/uuid"

type Citizenship struct {
	ID          int64     `json:"id"`
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	CountryUUID uuid.UUID `json:"country_uuid"`
}
