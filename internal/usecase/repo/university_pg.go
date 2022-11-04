package repo

import (
	"context"
	"fmt"
	"lg/internal/entity"
	"lg/internal/usecase"
	"lg/pkg/postgres"
)

type UniversityRepo struct {
	*postgres.Postgres
}

var _ usecase.UniversityRp = (*UniversityRepo)(nil)

func NewUniversityRepo(pg *postgres.Postgres) *UniversityRepo {
	return &UniversityRepo{pg}
}

func (c *UniversityRepo) GetAllUniversities(ctx context.Context) ([]entity.University, error) {
	query := `SELECT * FROM university`

	rows, err := c.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var universityList []entity.University
	for rows.Next() {
		university := entity.University{}
		err = rows.Scan(
			&university.ID,
			&university.UUID,
			&university.Name,
			&university.CityUUID,
		)
		if err != nil {
			return nil, fmt.Errorf("error in parsing project: %w", err)
		}
		universityList = append(universityList, university)
	}
	return universityList, nil
}
