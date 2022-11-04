package repo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
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
	query := `SELECT * FROM "user" WHERE email = $1`

	rows, err := u.Pool.Query(ctx, query, email)
	if err != nil {
		return entity.User{}, err
	}
	defer rows.Close()

	user := entity.User{}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.UUID, &user.Email, &user.Password)
		if err != nil {
			return entity.User{}, err
		}
	}
	return user, nil
}

func (u *UserRepo) StoreUser(ctx context.Context, user entity.User) error {
	query := `INSERT INTO "user" (email, password) VALUES($1, $2)`
	rows, err := u.Pool.Query(ctx, query, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("cannot insert value into users: %v", err)
	}
	defer rows.Close()
	return nil
}

func (u *UserRepo) GetUserByUUID(ctx context.Context, user uuid.UUID) (entity.User, error) {
	query := `SELECT * FROM "user" WHERE uuid = $1`

	rows, err := u.Pool.Query(ctx, query, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("cannot execute query: %v", err)
	}
	defer rows.Close()

	var usr entity.User
	for rows.Next() {
		err = rows.Scan(&usr.ID,
			&usr.UUID,
			&usr.Email,
			&usr.Password)
		if err != nil {
			return entity.User{}, err
		}
	}
	return usr, nil
}
