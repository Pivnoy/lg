package entity

type Project struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Link         string `json:"link"`
	Presentation string `json:"presentation"`
}
