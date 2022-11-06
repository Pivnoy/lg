package usecase

import (
	"context"
	"github.com/google/uuid"
)

type AchievementUseCase struct {
	repo AchievementRp
}

var _ AchievementContract = (*AchievementUseCase)(nil)

func NewAchievementUseCase(repo AchievementRp) *AchievementUseCase {
	return &AchievementUseCase{repo: repo}
}

func (a *AchievementUseCase) CreateAchievement(ctx context.Context, text string) (uuid.UUID, error) {
	return a.repo.CreateAchievement(ctx, text)
}
