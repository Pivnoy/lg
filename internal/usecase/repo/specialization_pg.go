package repo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"lg/internal/entity"
	"lg/internal/usecase"
	"lg/pkg/postgres"
)

type SpecializationRepo struct {
	*postgres.Postgres
}

var _ usecase.SpecializationRp = (*SpecializationRepo)(nil)

func NewSpecializationRepo(pg *postgres.Postgres) *SpecializationRepo {
	return &SpecializationRepo{pg}
}

func (c *SpecializationRepo) GetSpecializationByUUID(ctx context.Context, specializationKey uuid.UUID) (entity.Specialization, error) {
	query := `SELECT * FROM specialization where uuid=$1`

	rows, err := c.Pool.Query(ctx, query, specializationKey)
	if err != nil {
		return entity.Specialization{}, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	project := entity.Specialization{}
	for rows.Next() {
		err = rows.Scan(
			&project.ID,
			&project.UUID,
			&project.Name,
			&project.Value,
		)
		if err != nil {
			return entity.Specialization{}, fmt.Errorf("error in parsing specialization: %w", err)
		}
	}
	return project, nil
}

func (c *SpecializationRepo) GetAllSpecializations(ctx context.Context) ([]entity.Specialization, error) {
	query := `SELECT * FROM specialization`

	rows, err := c.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var specializationList []entity.Specialization
	for rows.Next() {
		specialization := entity.Specialization{}
		err = rows.Scan(
			&specialization.ID,
			&specialization.UUID,
			&specialization.Name,
			&specialization.Value,
		)
		if err != nil {
			return nil, fmt.Errorf("error in parsing project: %w", err)
		}
		specializationList = append(specializationList, specialization)
	}
	return specializationList, nil
}
