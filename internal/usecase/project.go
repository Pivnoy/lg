package usecase

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
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

func (p *ProjectUseCase) GetAllProjects(ctx context.Context) ([]entity.Project, error) {
	return p.repo.GetAllProjects(ctx)
}

func (p *ProjectUseCase) GetProjectByUUID(ctx context.Context, projectKey uuid.UUID) (entity.Project, error) {
	project, err := p.repo.GetProjectByUUID(ctx, projectKey)
	if (err == nil) && (project == entity.Project{}) {
		return project, fmt.Errorf("there is no project with this name")
	}
	return project, err
}

func (p *ProjectUseCase) CreateProject(ctx context.Context, project entity.Project) (uuid.UUID, error) {
	projectOld, err := p.repo.GetProjectByUUID(ctx, project.UUID)
	switch {
	case err != nil:
		return uuid.Nil, err
	case projectOld != entity.Project{}:
		return uuid.Nil, fmt.Errorf("project with that name already exists")
	}
	return p.repo.CreateProject(ctx, project)
}

func (p *ProjectUseCase) UpdateProjectByUUID(ctx context.Context, project entity.Project) error {
	_, err := p.GetProjectByUUID(ctx, project.UUID)
	if err != nil {
		return err
	}
	return p.repo.UpdateProjectByUUID(ctx, project)
}

func (p *ProjectUseCase) DeleteProjectByUUID(ctx context.Context, projectKey uuid.UUID) error {
	_, err := p.GetProjectByUUID(ctx, projectKey)
	if err != nil {
		return err
	}
	return p.repo.DeleteProjectByUUID(ctx, projectKey)
}
