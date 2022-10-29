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

func (p *ProjectRepo) GetProjectByName(ctx context.Context, name string) (entity.Project, error) {
	query := `SELECT * FROM project WHERE name = $1`

	rows, err := p.Pool.Query(ctx, query, name)
	if err != nil {
		return entity.Project{}, err
	}
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
