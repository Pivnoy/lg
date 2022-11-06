package repo

import (
	"context"
	"fmt"
	"lg/internal/entity"
	"lg/internal/usecase"
	"lg/pkg/postgres"
)

type RoleRepo struct {
	*postgres.Postgres
}

func NewRoleRepo(postgres *postgres.Postgres) *RoleRepo {
	return &RoleRepo{Postgres: postgres}
}

var _ usecase.RoleRp = (*RoleRepo)(nil)

func (r *RoleRepo) GetAllRoles(ctx context.Context) ([]entity.Role, error) {
	query := `SELECT * FROM role`

	rows, err := r.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var roleList []entity.Role
	for rows.Next() {
		role := entity.Role{}
		err = rows.Scan(
			&role.ID,
			&role.UUID,
			&role.Name,
			&role.SpecializationUUID,
		)
		if err != nil {
			return nil, fmt.Errorf("error in parsing category: %w", err)
		}
		roleList = append(roleList, role)
	}
	return roleList, nil
}
