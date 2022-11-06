package usecase

import (
	"context"
	"github.com/google/uuid"
	"lg/internal/entity"
)

type CityUseCase struct {
	repo CityRp
}

var _ CityContract = (*CityUseCase)(nil)

func NewCityUseCase(repo CityRp) *CityUseCase {
	return &CityUseCase{repo: repo}
}

func (c *CityUseCase) GetCitiesByCountryUUID(ctx context.Context, uuid uuid.UUID) ([]entity.City, error) {
	return c.repo.GetCitiesByCountryUUID(ctx, uuid)
}

func (c *CityUseCase) GetCityNameByUUID(ctx context.Context, cityKey uuid.UUID) (string, error) {
	return c.repo.GetCityNameByUUID(ctx, cityKey)
}
