package entity

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	ID         int64     `json:"id"`
	AuthorUUID uuid.UUID `json:"author_uuid"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}
