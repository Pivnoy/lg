package repo

import (
	"context"
	"fmt"
	"lg/internal/entity"
	"lg/internal/usecase"
	"lg/pkg/postgres"
)

type CitizenshipRepo struct {
	*postgres.Postgres
}

var _ usecase.CitizenshipRp = (*CitizenshipRepo)(nil)

func NewCitizenshipRepo(pg *postgres.Postgres) *CitizenshipRepo {
	return &CitizenshipRepo{pg}
}

func (c *CitizenshipRepo) GetAllCitizenships(ctx context.Context) ([]entity.Citizenship, error) {
	query := `SELECT * FROM citizenship`

	rows, err := c.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var citizenshipList []entity.Citizenship
	for rows.Next() {
		citizenship := entity.Citizenship{}
		err = rows.Scan(
			&citizenship.ID,
			&citizenship.UUID,
			&citizenship.Name,
			&citizenship.CountryUUID,
		)
		if err != nil {
			return nil, fmt.Errorf("error in parsing project: %w", err)
		}
		citizenshipList = append(citizenshipList, citizenship)
	}
	return citizenshipList, nil
}
