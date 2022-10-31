package entity

type Project struct {
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	ProjectLink      string `json:"link"`
	PresentationLink string `json:"presentation"`
	CreatorID        int64  `json:"creator_id"`
}
