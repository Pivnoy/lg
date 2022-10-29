package entity

type Profile struct {
	ID          int64  `json:"id"`
	Email       string `json:"email"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Country     string `json:"country"`
	City        string `json:"city"`
	Citizenship string `json:"citizenship"`
	Gender      string `json:"gender"`
	Contact     string `json:"contact"`
	Education   string `json:"education"`
	Employment  string `json:"employment"`
	Experience  uint64 `json:"experience"`
	Skill       string `json:"skill"`
	Achievement string `json:"achievement"`
	Team        string `json:"team"`
	TeamRole    string `json:"team_role"`
	Patent      string `json:"patent"`
	CompanyINN  uint64 `json:"company_inn"`
}
