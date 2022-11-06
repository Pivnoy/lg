package entity

import "github.com/google/uuid"

type Achievement struct {
	ID   int64     `json:"id"`
	UUID uuid.UUID `json:"uuid"`
	Text uuid.UUID `json:"text"`
}
