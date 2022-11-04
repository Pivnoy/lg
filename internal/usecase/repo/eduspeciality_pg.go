package repo

import (
	"context"
	"fmt"
	"lg/internal/entity"
	"lg/internal/usecase"
	"lg/pkg/postgres"
)

type EduspecialityRepo struct {
	*postgres.Postgres
}

var _ usecase.EduspecialityRp = (*EduspecialityRepo)(nil)

func NewEduspecialityRepo(pg *postgres.Postgres) *EduspecialityRepo {
	return &EduspecialityRepo{pg}
}

func (c *EduspecialityRepo) GetAllEduspecialities(ctx context.Context) ([]entity.Eduspeciality, error) {
	query := `SELECT * FROM eduspeciality`

	rows, err := c.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var eduspecialityList []entity.Eduspeciality
	for rows.Next() {
		eduspeciality := entity.Eduspeciality{}
		err = rows.Scan(
			&eduspeciality.ID,
			&eduspeciality.UUID,
			&eduspeciality.Name,
			&eduspeciality.Code,
		)
		if err != nil {
			return nil, fmt.Errorf("error in parsing project: %w", err)
		}
		eduspecialityList = append(eduspecialityList, eduspeciality)
	}
	return eduspecialityList, nil
}
