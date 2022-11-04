package entity

import "github.com/google/uuid"

type Country struct {
	ID   int64     `json:"id"`
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
	Code string    `json:"code"`
}
