package entity

import "github.com/google/uuid"

type University struct {
	ID       int64     `json:"id"`
	UUID     uuid.UUID `json:"uuid"`
	Name     string    `json:"name"`
	CityUUID uuid.UUID `json:"city_uuid"`
}
