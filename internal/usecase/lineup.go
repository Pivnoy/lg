package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"lg/internal/entity"
)

type LineupUseCase struct {
	repo LineupRp
}

func NewLineupUseCase(r LineupRp) *LineupUseCase {
	return &LineupUseCase{repo: r}
}

var _ LineupContract = (*LineupUseCase)(nil)

func (l *LineupUseCase) DeleteLineupByProjectUUID(ctx context.Context, projectKey uuid.UUID) error {
	exist, err := l.CheckLineupExistenceByProjectUUID(ctx, projectKey)
	switch {
	case err != nil:
		return err
	case exist:
		return l.repo.DeleteLineupByProjectUUID(ctx, projectKey)
	default:
		return fmt.Errorf("line up with project key %s missing", projectKey.String())
	}
}

func (l *LineupUseCase) CheckLineupExistenceByProjectUUID(ctx context.Context, projectKey uuid.UUID) (bool, error) {
	lineup, err := l.repo.GetLineupByProjectUUID(ctx, projectKey)
	switch {
	case err != nil:
		return false, err
	case err == nil && lineup == entity.Lineup{}:
		return false, nil
	default:
		return true, nil
	}
}
