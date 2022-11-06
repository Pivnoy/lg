package repo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
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

func (c *CitizenshipRepo) GetCitizenshipNameByUUID(ctx context.Context, citizenshipKey uuid.UUID) (string, error) {
	query := `SELECT name FROM citizenship WHERE uuid=$1`

	rows, err := c.Pool.Query(ctx, query, citizenshipKey)
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
