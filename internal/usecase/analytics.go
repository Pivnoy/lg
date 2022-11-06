package usecase

import (
	"context"
	"github.com/google/uuid"
	"lg/internal/entity"
)

type Analytics struct {
	repo AnalyticsRp
	pr   ProfileContract
}

var _ AnalyticsContract = (*Analytics)(nil)

func NewAnalytics(repo AnalyticsRp, pr ProfileContract) *Analytics {
	return &Analytics{repo: repo, pr: pr}
}

func (a *Analytics) Get(ctx context.Context, projUUID uuid.UUID) ([]entity.Profile, error) {
	uuids, err := a.repo.GetUserUUIDListOfVacsProject(ctx, projUUID)
	if err != nil {
		return nil, err
	}
	var profiles []entity.Profile
	for _, ud := range uuids {
		profile, err := a.pr.GetProfileByUser(ctx, ud)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}
	return profiles, nil
}
