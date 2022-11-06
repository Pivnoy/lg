package repo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"lg/internal/entity"
	"lg/internal/usecase"
	"lg/pkg/postgres"
)

type CountryRepo struct {
	*postgres.Postgres
}

var _ usecase.CountryRp = (*CountryRepo)(nil)

func NewCountryRepo(pg *postgres.Postgres) *CountryRepo {
	return &CountryRepo{pg}
}

func (c *CountryRepo) GetCountryNameByUUID(ctx context.Context, countryKey uuid.UUID) (string, error) {
	query := `SELECT name FROM country WHERE uuid=$1`

	rows, err := c.Pool.Query(ctx, query, countryKey)
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

func (c *CountryRepo) GetAllCountries(ctx context.Context) ([]entity.Country, error) {
	query := `SELECT * FROM country`

	rows, err := c.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var countryList []entity.Country
	for rows.Next() {
		country := entity.Country{}
		err = rows.Scan(
			&country.ID,
			&country.UUID,
			&country.Name,
			&country.Code,
		)
		if err != nil {
			return nil, fmt.Errorf("error in parsing category: %w", err)
		}
		countryList = append(countryList, country)
	}
	return countryList, nil
}
