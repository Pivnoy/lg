package entity

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	ID           int64     `json:"id"`
	AuthorUUID   uuid.UUID `json:"author_uuid"`
	Type         string    `json:"msg_type"`
	Content      string    `json:"content"`
	CreationDate time.Time `json:"creation_date"`
	ChatUUID     uuid.UUID `json:"chat_uuid"`
}
