package repo

import (
	"context"
	"fmt"
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
