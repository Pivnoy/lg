package usecase

import (
	"context"
	"lg/internal/entity"
)

type SpecializationUseCase struct {
	repo SpecializationRp
}

var _ SpecializationContract = (*SpecializationUseCase)(nil)

func NewSpecializationUseCase(repo SpecializationRp) *SpecializationUseCase {
	return &SpecializationUseCase{repo: repo}
}

func (c *SpecializationUseCase) GetAllSpecializations(ctx context.Context) ([]entity.Specialization, error) {
	return c.repo.GetAllSpecializations(ctx)
}
