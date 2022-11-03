package usecase

import (
	"context"
	"github.com/google/uuid"
	"lg/internal/entity"
)

type ProfileUseCase struct {
	repo ProfileRp
}

var _ ProfileContract = (*ProfileUseCase)(nil)

func NewProfileUseCase(repo ProfileRp) *ProfileUseCase {
	return &ProfileUseCase{
		repo: repo,
	}
}

func (p *ProfileUseCase) GetProfileByUser(ctx context.Context, uuid uuid.UUID) (entity.Profile, error) {
	return p.repo.GetProfileByUser(ctx, uuid)
}
