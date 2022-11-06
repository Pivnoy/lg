package repo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"lg/internal/usecase"
	"lg/pkg/postgres"
)

type AchievementRepo struct {
	*postgres.Postgres
}

var _ usecase.AchievementRp = (*AchievementRepo)(nil)

func NewAchievementRepo(pg *postgres.Postgres) *AchievementRepo {
	return &AchievementRepo{pg}
}

func (a *AchievementRepo) CreateAchievement(ctx context.Context, text string) (uuid.UUID, error) {
	query := `INSERT INTO achievement (text) VALUES ($1) RETURNING uuid`
	rows, err := a.Pool.Query(ctx, query, text)
	if err != nil {
		return uuid.Nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var achievementKey uuid.UUID
	for rows.Next() {
		err = rows.Scan(&achievementKey)
		if err != nil {
			return uuid.Nil, fmt.Errorf("error in parsing uuid achievement: %w", err)
		}
	}
	return achievementKey, nil
}
