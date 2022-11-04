package usecase

import (
	"context"
	"lg/internal/entity"
)

type EmploymentUseCase struct {
	repo EmploymentRp
}

var _ EmploymentContract = (*EmploymentUseCase)(nil)

func NewEmploymentUseCase(repo EmploymentRp) *EmploymentUseCase {
	return &EmploymentUseCase{repo: repo}
}

func (c *EmploymentUseCase) GetAllEmployments(ctx context.Context) ([]entity.Employment, error) {
	return c.repo.GetAllEmployments(ctx)
}
