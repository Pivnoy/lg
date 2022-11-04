package repo

import (
	"context"
	"fmt"
	"lg/internal/entity"
	"lg/internal/usecase"
	"lg/pkg/postgres"
)

type EmploymentRepo struct {
	*postgres.Postgres
}

var _ usecase.EmploymentRp = (*EmploymentRepo)(nil)

func NewEmploymentRepo(pg *postgres.Postgres) *EmploymentRepo {
	return &EmploymentRepo{pg}
}

func (c *EmploymentRepo) GetAllEmployments(ctx context.Context) ([]entity.Employment, error) {
	query := `SELECT * FROM employment`

	rows, err := c.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var employmentList []entity.Employment
	for rows.Next() {
		employment := entity.Employment{}
		err = rows.Scan(
			&employment.ID,
			&employment.UUID,
			&employment.Name,
			&employment.Value,
		)
		if err != nil {
			return nil, fmt.Errorf("error in parsing project: %w", err)
		}
		employmentList = append(employmentList, employment)
	}
	return employmentList, nil
}