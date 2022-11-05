package repo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"lg/internal/entity"
	"lg/internal/usecase"
	"lg/pkg/postgres"
	"time"
)

type ProfileRepo struct {
	*postgres.Postgres
}

var _ usecase.ProfileRp = (*ProfileRepo)(nil)

func NewProfileRepo(pg *postgres.Postgres) *ProfileRepo {
	return &ProfileRepo{pg}
}

func check(val uuid.UUID) any {
	if val != uuid.Nil {
		return val
	}
	return nil
}

func (p *ProfileRepo) CheckFkProfile(ctx context.Context, profile entity.Profile) (string, error) {
	query := `SELECT check_fk_profile($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	rows, err := p.Pool.Query(
		ctx,
		query,
		profile.UserUUID,
		profile.CountryUUID,
		profile.CityUUID,
		profile.CitizenshipUUID,
		profile.UniversityUUID,
		profile.EduspecialityUUID,
		profile.EmploymentUUID,
		profile.AchievementUUID,
		profile.TeamUUID,
		profile.SpecializationUUID,
		profile.CompanyUUID)
	if err != nil {
		return "", fmt.Errorf("cannot execute check fk profile function, query: %v", err)
	}
	defer rows.Close()
	var res string
	for rows.Next() {
		err = rows.Scan(&res)
		if err != nil {
			return "", fmt.Errorf("error in parsing result of check function profile: %w", err)
		}
	}
	return res, nil
}

func (p *ProfileRepo) CreateProfile(ctx context.Context, profile entity.Profile) (entity.Profile, error) {
	query := `INSERT INTO profile (
                     user_uuid, 
                     firstname, 
                     lastname, 
                     patronymic, 
                     country_uuid, 
                     city_uuid, 
                     citizenship_uuid, 
                     gender, 
                     phone, 
                     email, 
                     university_uuid, 
                     eduspeciality_uuid, 
                     graduation_year, 
                     employment_uuid, 
                     experience, 
                     achievement_uuid, 
                     team_uuid, 
                     specialization_uuid, 
                     company_uuid,
                     creation_date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)`
	newDate := time.Now()
	//new_uni := check(profile.UniversityUUID)
	//new_edu := check(profile.EduspecialityUUID)
	//new_team := check(profile.TeamUUID)
	//new_com := check(profile.CompanyUUID)
	rows, err := p.Pool.Query(ctx,
		query,
		profile.UserUUID,
		profile.Firstname,
		profile.Lastname,
		profile.Patronymic,
		profile.CountryUUID,
		profile.CityUUID,
		profile.CitizenshipUUID,
		profile.Gender,
		profile.Phone,
		profile.Email,
		check(profile.UniversityUUID),    // NULL
		check(profile.EduspecialityUUID), // NULL
		profile.GraduationYear,
		profile.EmploymentUUID,
		profile.Experience,
		profile.AchievementUUID,
		check(profile.TeamUUID), // NULL
		profile.SpecializationUUID,
		check(profile.CompanyUUID), // NULL
		newDate,
	)
	if err != nil {
		return entity.Profile{}, fmt.Errorf("cannot execute profile query: %v", err)
	}
	defer rows.Close()
	return profile, nil
}

func (p *ProfileRepo) GetProfileByUser(ctx context.Context, user uuid.UUID) (entity.Profile, error) {
	query := `SELECT * FROM profile WHERE user_uuid = $1`

	rows, err := p.Pool.Query(ctx, query, user)
	if err != nil {
		return entity.Profile{}, fmt.Errorf("cannot execute profile query: %v", err)
	}

	var profile entity.Profile
	for rows.Next() {
		//&profile.ID,
		err = rows.Scan(
			&profile.UserUUID,
			&profile.Firstname,
			&profile.Lastname,
			&profile.Patronymic,
			&profile.CountryUUID,
			&profile.CityUUID,
			&profile.CitizenshipUUID,
			&profile.Gender,
			&profile.Phone,
			&profile.Email,
			&profile.UniversityUUID,
			&profile.EduspecialityUUID,
			&profile.GraduationYear,
			&profile.EmploymentUUID,
			&profile.Experience,
			&profile.AchievementUUID,
			&profile.TeamUUID,
			&profile.SpecializationUUID,
			&profile.CompanyUUID)
		if err != nil {
			return entity.Profile{}, fmt.Errorf("cannot parse profile: %v", err)
		}
	}
	return profile, nil
}
