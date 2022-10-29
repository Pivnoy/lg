package usecase

import (
	"context"
	"lg/internal/entity"
)

type UserUseCase struct {
	repo UserRp
}

var _ UserContract = (*UserUseCase)(nil)

func NewUserUseCase(repo UserRp) *UserUseCase {
	return &UserUseCase{
		repo: repo,
	}
}

func (u *UserUseCase) GetUser(ctx context.Context, email string) (entity.User, error) {
	return u.repo.GetUserByEmail(ctx, email)
}

func (u *UserUseCase) StoreUser(ctx context.Context, user entity.User) error {
	return u.repo.StoreUser(ctx, user)
}
