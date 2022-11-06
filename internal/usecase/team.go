package usecase

import (
	"context"
	"github.com/google/uuid"
	"lg/internal/entity"
)

type TeamUseCase struct {
	repo TeamRp
}

var _ TeamContract = (*TeamUseCase)(nil)

func NewTeamUseCase(repo TeamRp) *TeamUseCase {
	return &TeamUseCase{repo: repo}
}

func (t *TeamUseCase) CreateTeam(ctx context.Context, team entity.Team) (uuid.UUID, error) {
	return t.repo.CreateTeam(ctx, team)
}
