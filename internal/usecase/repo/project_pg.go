package repo

import (
	"context"
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
		return nil, err
	}
	defer rows.Close()
	var projectList []entity.Project
	for rows.Next() {
		project := entity.Project{}
		err = rows.Scan(
			&project.ID,
			&project.Name,
			&project.Description,
			&project.Link,
			&project.Presentation,
		)
		if err != nil {
			return nil, err
		}
	}
	return projectList, nil
}

func (p *ProjectRepo) GetProjectByName(ctx context.Context, name string) (entity.Project, error) {
	query := `SELECT * FROM project WHERE name = $1`

	rows, err := p.Pool.Query(ctx, query, name)
	if err != nil {
		return entity.Project{}, err
	}
	defer rows.Close()
	project := entity.Project{}
	for rows.Next() {
		err = rows.Scan(
			&project.ID,
			&project.Name,
			&project.Description,
			&project.Link,
			&project.Presentation,
		)
		if err != nil {
			return entity.Project{}, err
		}
	}
	return project, nil
}

func (p *ProjectRepo) CreateProject(ctx context.Context, project entity.Project) (string, error) {
	query := `INSERT INTO project (id, name, description, link, presentation) VALUES ($1, $2, $3, $4, $5) RETURNING name`

	rows, err := p.Pool.Query(ctx, query, project.ID, project.Name, project.Description, project.Link, project.Presentation)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			return "", err
		}
	}
	return name, nil
}
