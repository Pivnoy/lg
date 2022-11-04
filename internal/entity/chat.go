package entity

import "github.com/google/uuid"

// Chat Add here project UUID
type Chat struct {
	ID   int64     `json:"id"`
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}
