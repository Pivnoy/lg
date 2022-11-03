package repo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"lg/internal/entity"
	"lg/internal/usecase"
	"lg/pkg/postgres"
)

type ProfileRepo struct {
	*postgres.Postgres
}

var _ usecase.ProfileRp = (*ProfileRepo)(nil)

func NewProfileRepo(pg *postgres.Postgres) *ProfileRepo {
	return &ProfileRepo{pg}
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
			&profile.CompanyUUID,
			&profile.UUID)
		if err != nil {
			return entity.Profile{}, fmt.Errorf("cannot parse profile: %v", err)
		}
	}
	return profile, nil
}
