package repo

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"lg/internal/entity"
	"lg/internal/usecase"
	"lg/pkg/postgres"
)

type ProjectRepo struct {
	*postgres.Postgres
}

var _ usecase.ProjectRp = (*ProjectRepo)(nil)

func NewProjectRepo(pg *postgres.Postgres) *ProjectRepo {
	return &ProjectRepo{pg}
}

func (p *ProjectRepo) GetAllProjects(ctx context.Context) ([]entity.Project, error) {
	query := `SELECT * FROM project`

	rows, err := p.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var projectList []entity.Project
	for rows.Next() {
		project := entity.Project{}
		err = rows.Scan(
			&project.ID,
			&project.UUID,
			&project.Name,
			&project.Description,
			&project.ProjectLink,
			&project.PresentationLink,
			&project.CreatorID,
		)
		if err != nil {
			return nil, fmt.Errorf("error in parsing project: %w", err)
		}
		projectList = append(projectList, project)
	}
	return projectList, nil
}

func (p *ProjectRepo) GetProjectByUUID(ctx context.Context, projectKey uuid.UUID) (entity.Project, error) {
	query := `SELECT * FROM project WHERE uuid = $1`

	rows, err := p.Pool.Query(ctx, query, projectKey)
	if err != nil {
		return entity.Project{}, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	project := entity.Project{}
	for rows.Next() {
		err = rows.Scan(
			&project.ID,
			&project.UUID,
			&project.Name,
			&project.Description,
			&project.ProjectLink,
			&project.PresentationLink,
			&project.CreatorID,
		)
		if err != nil {
			return entity.Project{}, fmt.Errorf("error in parsing project: %w", err)
		}
	}
	return project, nil
}

func (p *ProjectRepo) CreateProject(ctx context.Context, project entity.Project) (uuid.UUID, error) {
	query := `INSERT INTO project (name, description, project_link, presentation_link, creator_id) VALUES ($1, $2, $3, $4, $5) RETURNING uuid`

	rows, err := p.Pool.Query(ctx, query, project.Name, project.Description, project.ProjectLink, project.PresentationLink, project.CreatorID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var projectKey uuid.UUID
	for rows.Next() {
		err = rows.Scan(&projectKey)
		if err != nil {
			return uuid.Nil, fmt.Errorf("error in parsing name of project: %w", err)
		}
	}
	return projectKey, nil
}

func (p *ProjectRepo) UpdateProjectByUUID(ctx context.Context, project entity.Project) error {
	query := `UPDATE project SET description=$1, project_link=$2, presentation_link=$3, creator_id=$4 where uuid = $5`

	rows, err := p.Pool.Query(ctx, query, project.Description, project.ProjectLink, project.PresentationLink, project.CreatorID, project.UUID)
	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	return nil
}

func (p *ProjectRepo) DeleteProjectByUUID(ctx context.Context, projectKey uuid.UUID) error {
	query := `DELETE FROM project WHERE uuid=$1`

	rows, err := p.Pool.Query(ctx, query, projectKey)
	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	return nil
}
