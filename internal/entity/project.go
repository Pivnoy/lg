package entity

import (
	"github.com/google/uuid"
	"time"
)

type Project struct {
	ID               int64     `json:"id"`
	UUID             uuid.UUID `json:"uuid"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	CategoryUUID     uuid.UUID `json:"categoryUUID"`
	ProjectLink      string    `json:"link"`
	PresentationLink string    `json:"presentation"`
	CreatorUUID      uuid.UUID `json:"creatorUUID"`
	IsVisible        string    `json:"isVisible"`
	CreationDate     time.Time `json:"creationDate"`
}
