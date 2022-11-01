package usecase

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"lg/internal/entity"
)

type RegisterUseCase struct {
	u UserContract
}

var _ RegisterContract = (*RegisterUseCase)(nil)

func NewSignInUseCase(u UserContract) *RegisterUseCase {
	return &RegisterUseCase{
		u: u,
	}
}

func (s *RegisterUseCase) CreateNewUser(ctx context.Context, email string, password string) error {
	if len(email) <= 4 || len(password) <= 4 {
		return fmt.Errorf("invalid format of username or password")
	}
	exists, err := s.u.CheckUserExistence(ctx, email)
	if err != nil {
		return fmt.Errorf("error in checking user existence: %v", err)
	}
	if !exists {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
		if err != nil {
			return fmt.Errorf("error in hashing psswrd: %v", err)
		}
		us := entity.User{
			Email:    email,
			Password: string(passwordHash),
		}
		return s.u.StoreUser(ctx, us)
	}
	return fmt.Errorf("user already exists")
}
