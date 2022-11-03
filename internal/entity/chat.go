package entity

import "github.com/google/uuid"

type Chat struct {
	ID       int64       `json:"id"`
	UUID     uuid.UUID   `json:"uuid"`
	Users    []uuid.UUID `json:"users"`
	Messages []uuid.UUID `json:"messages"`
}
