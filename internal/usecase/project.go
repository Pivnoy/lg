package usecase

import (
	"context"
	"fmt"
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

func (p *ProjectUseCase) GetProjectByName(ctx context.Context, name string) (entity.Project, error) {
	project, err := p.repo.GetProjectByName(ctx, name)
	if (err == nil) && (project == entity.Project{}) {
		return project, fmt.Errorf("there is no project with this name")
	}
	return project, err
}

func (p *ProjectUseCase) CreateProject(ctx context.Context, project entity.Project) (string, error) {
	projectOld, err := p.repo.GetProjectByName(ctx, project.Name)
	switch {
	case err != nil:
		return "", err
	case projectOld != entity.Project{}:
		return "", fmt.Errorf("project with that name already exists")
	}
	return p.repo.CreateProject(ctx, project)
}

func (p *ProjectUseCase) UpdateProject(ctx context.Context, project entity.Project) error {
	_, err := p.GetProjectByName(ctx, project.Name)
	if err != nil {
		return err
	}
	return p.repo.UpdateProject(ctx, project)
}

func (p *ProjectUseCase) DeleteProject(ctx context.Context, name string) error {
	_, err := p.GetProjectByName(ctx, name)
	if err != nil {
		return err
	}
	return p.repo.DeleteProject(ctx, name)
}
