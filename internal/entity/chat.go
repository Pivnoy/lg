package entity

import "github.com/google/uuid"

type Chat struct {
	ID          int64     `json:"id"`
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	ProjectUUID string    `json:"project_uuid"`
}
