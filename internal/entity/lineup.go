package entity

import "github.com/google/uuid"

type Lineup struct {
	Id          int64     `json:"id"`
	UUID        uuid.UUID `json:"uuid"`
	TeamUUID    uuid.UUID `json:"team_uuid"`
	RoleUUID    uuid.UUID `json:"role_uuid"`
	ProfileUUID uuid.UUID `json:"profile_uuid"`
	ProjectUUID uuid.UUID `json:"project_uuid"`
}
