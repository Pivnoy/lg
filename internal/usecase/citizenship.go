package usecase

import (
	"context"
	"lg/internal/entity"
)

type CitizenshipUseCase struct {
	repo CitizenshipRp
}

var _ CitizenshipContract = (*CitizenshipUseCase)(nil)

func NewCitizenshipUseCase(repo CitizenshipRp) *CitizenshipUseCase {
	return &CitizenshipUseCase{repo: repo}
}

func (c *CitizenshipUseCase) GetAllCitizenships(ctx context.Context) ([]entity.Citizenship, error) {
	return c.repo.GetAllCitizenships(ctx)
}
