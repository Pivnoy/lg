package entity

import "github.com/google/uuid"

type Profile struct {
	ID                 uuid.UUID `json:"id"`
	UserUUID           uuid.UUID `json:"user_uuid"`
	Firstname          string    `json:"firstname"`
	Lastname           string    `json:"lastname"`
	Patronymic         string    `json:"patronymic"`
	CountryUUID        uuid.UUID `json:"country_uuid"`
	CityUUID           uuid.UUID `json:"city_uuid"`
	CitizenshipUUID    uuid.UUID `json:"citizenship_uuid"`
	Gender             string    `json:"gender"`
	Phone              string    `json:"phone"`
	Email              string    `json:"email"`
	UniversityUUID     uuid.UUID `json:"university_uuid"`
	EduspecialityUUID  uuid.UUID `json:"eduspeciality_uuid"`
	GraduationYear     uint      `json:"graduation_year"`
	EmploymentUUID     uuid.UUID `json:"employment_uuid"`
	Experience         uint      `json:"experience"`
	AchievementUUID    uuid.UUID `json:"achievement_uuid"`
	TeamUUID           uuid.UUID `json:"team_uuid"`
	SpecializationUUID uuid.UUID `json:"specialization_uuid"`
	CompanyUUID        uuid.UUID `json:"company_uuid"`
	UUID               uuid.UUID `json:"uuid"`
}
