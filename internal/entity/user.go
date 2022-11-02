package entity

import "github.com/google/uuid"

type User struct {
	ID       int64     `json:"id"`
	UUID     uuid.UUID `json:"uuid"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
