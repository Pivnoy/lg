package repo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
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

func (c *UniversityRepo) GetNameUniversityByUUID(ctx context.Context, universityKey uuid.UUID) (string, error) {
	query := `SELECT name FROM university WHERE uuid=$1`

	rows, err := c.Pool.Query(ctx, query, universityKey)
	if err != nil {
		return "", fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var name string
	for rows.Next() {
		err = rows.Scan(
			&name,
		)
		if err != nil {
			return "", fmt.Errorf("error in parsing category: %w", err)
		}
	}
	return name, nil
}
