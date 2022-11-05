package usecase

import (
	"context"
	"lg/internal/entity"
)

type CategoryUseCase struct {
	repo CategoryRp
}

var _ CategoryContract = (*CategoryUseCase)(nil)

func NewCategoryUseCase(repo CategoryRp) *CategoryUseCase {
	return &CategoryUseCase{repo: repo}
}

func (c *CategoryUseCase) GetAllCategory(ctx context.Context) ([]entity.Category, error) {
	return c.repo.GetAllCategory(ctx)
}
