package usecase

import (
	"context"
	"github.com/google/uuid"
	"lg/internal/entity"
)

type EduspecialityUseCase struct {
	repo EduspecialityRp
}

var _ EduspecialityContract = (*EduspecialityUseCase)(nil)

func NewEduspecialityUseCase(repo EduspecialityRp) *EduspecialityUseCase {
	return &EduspecialityUseCase{repo: repo}
}

func (c *EduspecialityUseCase) GetAllEduspecialities(ctx context.Context) ([]entity.Eduspeciality, error) {
	return c.repo.GetAllEduspecialities(ctx)
}

func (c *EduspecialityUseCase) GetEduspecialityNameByUUID(ctx context.Context, eduspecialityKey uuid.UUID) (string, error) {
	return c.repo.GetEduspecialityNameByUUID(ctx, eduspecialityKey)
}
