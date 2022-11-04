package repo

import (
	"context"
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

}
