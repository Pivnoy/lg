package repo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"lg/internal/entity"
	"lg/internal/usecase"
	"lg/pkg/postgres"
)

type LineupRepo struct {
	*postgres.Postgres
}

func NewLineupRepo(pg *postgres.Postgres) *LineupRepo {
	return &LineupRepo{pg}
}

var _ usecase.LineupRp = (*LineupRepo)(nil)

func (l *LineupRepo) CreateLineup(ctx context.Context, lineup entity.Lineup) error {
	query := `INSERT INTO lineup (team_uuid, role_uuid, profile_uuid, project_uuid) VALUES ($1, $2, $3, $4)`
	rows, err := l.Pool.Query(ctx, query, lineup.TeamUUID, lineup.RoleUUID, nil, lineup.ProjectUUID)
	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	return nil
}

func (l *LineupRepo) GetLineupByProjectUUID(ctx context.Context, projectKey uuid.UUID) (entity.Lineup, error) {
	query := `SELECT * FROM lineup WHERE project_uuid = $1`

	rows, err := l.Pool.Query(ctx, query, projectKey)
	if err != nil {
		return entity.Lineup{}, err
	}
	defer rows.Close()

	lineup := entity.Lineup{}
	for rows.Next() {
		err = rows.Scan(
			&lineup.Id,
			&lineup.UUID,
			&lineup.TeamUUID,
			&lineup.RoleUUID,
			&lineup.ProfileUUID,
			&lineup.ProjectUUID,
		)
		if err != nil {
			return entity.Lineup{}, err
		}
	}
	return lineup, nil
}

func (l *LineupRepo) DeleteLineupByProjectUUID(ctx context.Context, projectKey uuid.UUID) error {
	query := `DELETE FROM lineup WHERE project_uuid=$1`

	rows, err := l.Pool.Query(ctx, query, projectKey)
	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	return nil
}
