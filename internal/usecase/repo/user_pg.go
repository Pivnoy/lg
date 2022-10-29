package repo

import (
	"context"
	"fmt"
	"lg/internal/entity"
	"lg/internal/usecase"
	"lg/pkg/postgres"
)

type UserRepo struct {
	*postgres.Postgres
}

var _ usecase.UserRp = (*UserRepo)(nil)

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

func (u *UserRepo) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	query := `SELECT * FROM users WHERE email = $1`

	rows, err := u.Pool.Query(ctx, query, email)
	if err != nil {
		return entity.User{}, err
	}
	defer rows.Close()

	user := entity.User{}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return entity.User{}, err
		}
	}
	return user, nil
}

func (u *UserRepo) StoreUser(ctx context.Context, user entity.User) error {
	query := `INSERT INTO users (email, password) VALUES($1, $2)`
	rows, err := u.Pool.Query(ctx, query, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("cannot insert value into users: %v", err)
	}
	defer rows.Close()
	return nil
}
