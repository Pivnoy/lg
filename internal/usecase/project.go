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
	t    TeamContract
}

var _ ProjectContract = (*ProjectUseCase)(nil)

func NewProjectUseCase(repo ProjectRp, l LineupContract, t TeamContract) *ProjectUseCase {
	return &ProjectUseCase{repo: repo, l: l, t: t}
}

func (p *ProjectUseCase) ChangeVisibility(ctx context.Context, projectKey uuid.UUID) error {
	visibility, err := p.GetVisibilityByProject(ctx, projectKey)
	if err != nil {
		return err
	}
	switch visibility {
	case "visible":
		visibility = "invisible"
	case "invisible":
		visibility = "visible"
	default:
		return fmt.Errorf("invalid visiblity")
	}
	return p.repo.ChangeVisibility(ctx, projectKey, visibility)
}

func (p *ProjectUseCase) GetVisibilityByProject(ctx context.Context, projectKey uuid.UUID) (string, error) {
	return p.repo.GetVisibilityByProject(ctx, projectKey)
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

func (p *ProjectUseCase) CreateProject(ctx context.Context, project entity.Project, roleKeys []uuid.UUID) (uuid.UUID, error) {
	// создание проекта
	projectKey, err := p.repo.CreateProject(ctx, project)
	if err != nil {
		return projectKey, err
	}
	// создание команды
	teamKey, err := p.t.CreateTeam(ctx, entity.Team{
		UUID:  projectKey,
		Name:  project.Name,
		Value: project.Name,
	})
	if err != nil {
		return uuid.Nil, err
	}
	// создание лайнапов
	for _, v := range roleKeys {
		err = p.l.CreateLineup(ctx, entity.Lineup{
			TeamUUID:    teamKey,
			RoleUUID:    v,
			ProfileUUID: uuid.Nil,
			ProjectUUID: projectKey,
		})
		if err != nil {
			return uuid.Nil, err
		}
	}
	return projectKey, nil
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
