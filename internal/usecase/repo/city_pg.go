package repo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"lg/internal/entity"
	"lg/internal/usecase"
	"lg/pkg/postgres"
)

type CityRepo struct {
	*postgres.Postgres
}

var _ usecase.CityRp = (*CityRepo)(nil)

func NewCityRepo(pg *postgres.Postgres) *CityRepo {
	return &CityRepo{pg}
}

func (c *CityRepo) GetCitiesByCountryUUID(ctx context.Context, countryKey uuid.UUID) ([]entity.City, error) {
	query := `SELECT * FROM city where country_uuid=$1`

	rows, err := c.Pool.Query(ctx, query, countryKey)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var citiesList []entity.City
	for rows.Next() {
		city := entity.City{}
		err = rows.Scan(
			&city.ID,
			&city.UUID,
			&city.Name,
			&city.CountryUUID,
		)
		if err != nil {
			return nil, fmt.Errorf("error in parsing city: %w", err)
		}
		citiesList = append(citiesList, city)
	}
	return citiesList, nil
}
