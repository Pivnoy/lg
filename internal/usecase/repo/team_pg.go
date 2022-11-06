package repo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"lg/internal/entity"
	"lg/internal/usecase"
	"lg/pkg/postgres"
)

type TeamRepo struct {
	*postgres.Postgres
}

var _ usecase.TeamRp = (*TeamRepo)(nil)

func NewTeamRepo(postgres *postgres.Postgres) *TeamRepo {
	return &TeamRepo{Postgres: postgres}
}

func (t *TeamRepo) CreateTeam(ctx context.Context, team entity.Team) (uuid.UUID, error) {
	query := `INSERT INTO team (uuid, name, value) VALUES ($1, $2, $3) RETURNING uuid`
	rows, err := t.Pool.Query(ctx, query, team.UUID, team.Name, team.Name)
	if err != nil {
		return uuid.Nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var teamKey uuid.UUID
	for rows.Next() {
		err = rows.Scan(&teamKey)
		if err != nil {
			return uuid.Nil, fmt.Errorf("error in parsing uuid of team: %w", err)
		}
	}
	return teamKey, nil
}
