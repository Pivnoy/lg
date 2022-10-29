package usecase

import (
	"context"
	"lg/internal/entity"
)

type ProjectUseCase struct {
	repo ProjectRp
}

var _ ProjectContract = (*ProjectUseCase)(nil)

func NewProjectUseCase(repo ProjectRp) *ProjectUseCase {
	return &ProjectUseCase{
		repo: repo,
	}
}

func (p *ProjectUseCase) GetProjectByName(ctx context.Context, name string) (entity.Project, error) {
	return p.repo.GetProjectByName(ctx, name)
}
