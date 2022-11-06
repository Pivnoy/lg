package usecase

import (
	"context"
	"lg/internal/entity"
)

type RoleUseCase struct {
	repo RoleRp
}

func NewRoleUseCase(repo RoleRp) *RoleUseCase {
	return &RoleUseCase{repo: repo}
}

var _ RoleContract = (*RoleUseCase)(nil)

func (r *RoleUseCase) GetAllRoles(ctx context.Context) ([]entity.Role, error) {
	return r.repo.GetAllRoles(ctx)
}
