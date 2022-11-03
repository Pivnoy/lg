package usecase

import (
	"context"
	"lg/internal/entity"
)

type UniversityUseCase struct {
	repo UniversityRp
}

var _ UniversityContract = (*UniversityUseCase)(nil)

func NewUniversityUseCase(repo UniversityRp) *UniversityUseCase {
	return &UniversityUseCase{repo: repo}
}

func (c *UniversityUseCase) GetAllUniversities(ctx context.Context) ([]entity.University, error) {
	return c.repo.GetAllUniversities(ctx)
}
