package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"lg/internal/entity"
)

type ProjectUseCase struct {
	repo ProjectRp
	l    LineupContract
}

var _ ProjectContract = (*ProjectUseCase)(nil)

func NewProjectUseCase(repo ProjectRp, l LineupContract) *ProjectUseCase {
	return &ProjectUseCase{
		repo: repo,
		l:    l,
	}
}

func (p *ProjectUseCase) GetAllProjects(ctx context.Context, page, limit uint) ([]entity.Project, error) {
	return p.repo.GetAllProjects(ctx, page, limit)
}

func (p *ProjectUseCase) GetProjectByUUID(ctx context.Context, projectKey uuid.UUID) (entity.Project, error) {
	project, err := p.repo.GetProjectByUUID(ctx, projectKey)
	if (err == nil) && (project == entity.Project{}) {
		return project, fmt.Errorf("there is no project with this name")
	}
	return project, err
}

func (p *ProjectUseCase) CreateProject(ctx context.Context, project entity.Project) (uuid.UUID, error) {
	return p.repo.CreateProject(ctx, project)
}

func (p *ProjectUseCase) DeleteProjectByUUID(ctx context.Context, projectKey uuid.UUID) error {
	err := p.l.DeleteLineupByProjectUUID(ctx, projectKey)
	if err != nil {
		return err
	}
	exist, err := p.CheckProjectExistenceByProjectUUID(ctx, projectKey)
	switch {
	case err != nil:
		return err
	case exist:
		return p.repo.DeleteProjectByUUID(ctx, projectKey)
	default:
		return fmt.Errorf("project with project key %s missing", projectKey.String())
	}
}

func (p *ProjectUseCase) CheckProjectExistenceByProjectUUID(ctx context.Context, projectKey uuid.UUID) (bool, error) {
	project, err := p.repo.GetProjectByUUID(ctx, projectKey)
	switch {
	case err != nil:
		return false, err
	case err == nil && project == entity.Project{}:
		return false, nil
	default:
		return true, nil
	}
}
