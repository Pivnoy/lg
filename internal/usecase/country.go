package usecase

import (
	"context"
	"lg/internal/entity"
)

type CountryUseCase struct {
	repo CountryRp
}

var _ CountryContract = (*CountryUseCase)(nil)

func NewCountryUseCase(repo CountryRp) *CountryUseCase {
	return &CountryUseCase{repo: repo}
}

func (c *CountryUseCase) GetAllCountries(ctx context.Context) ([]entity.Country, error) {
	return c.repo.GetAllCountries(ctx)
}
