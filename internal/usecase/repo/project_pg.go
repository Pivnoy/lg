package repo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"lg/internal/entity"
	"lg/internal/usecase"
	"lg/pkg/postgres"
	"time"
)

type ProjectRepo struct {
	*postgres.Postgres
}

var _ usecase.ProjectRp = (*ProjectRepo)(nil)

func NewProjectRepo(pg *postgres.Postgres) *ProjectRepo {
	return &ProjectRepo{pg}
}

func (p *ProjectRepo) GetAllProjects(ctx context.Context, page, limit uint) ([]entity.Project, error) {

	offset := (page - 1) * limit

	query := `SELECT * FROM project order by id offset $1 limit $2`

	rows, err := p.Pool.Query(ctx, query, offset, limit)
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
			&project.CategoryUUID,
			&project.ProjectLink,
			&project.PresentationLink,
			&project.CreatorUUID,
			&project.IsVisible,
			&project.CreationDate,
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
			&project.CategoryUUID,
			&project.ProjectLink,
			&project.PresentationLink,
			&project.CreatorUUID,
			&project.IsVisible,
			&project.CreationDate,
		)
		if err != nil {
			return entity.Project{}, fmt.Errorf("error in parsing project: %w", err)
		}
	}
	return project, nil
}

func (p *ProjectRepo) CreateProject(ctx context.Context, project entity.Project) (uuid.UUID, error) {
	query := `INSERT INTO project (name, description, category_uuid, project_link, presentation_link, creator_uuid, is_visible, creation_date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING uuid`
	newDate := time.Now()
	rows, err := p.Pool.Query(ctx, query, project.Name, project.Description, project.CategoryUUID, project.ProjectLink, project.PresentationLink, project.CreatorUUID, project.IsVisible, newDate)
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

func (p *ProjectRepo) DeleteProjectByUUID(ctx context.Context, projectKey uuid.UUID) error {
	query := `DELETE FROM project WHERE uuid=$1`

	rows, err := p.Pool.Query(ctx, query, projectKey)
	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	return nil
}
