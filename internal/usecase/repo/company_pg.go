package repo

import (
	"context"
	"fmt"
	"lg/internal/entity"
	"lg/internal/usecase"
	"lg/pkg/postgres"
)

type CompanyRepo struct {
	*postgres.Postgres
}

var _ usecase.CompanyRp = (*CompanyRepo)(nil)

func NewCompanyRepo(pg *postgres.Postgres) *CompanyRepo {
	return &CompanyRepo{pg}
}

func (c *CompanyRepo) GetCompanyByInn(ctx context.Context, inn string) (entity.Company, error) {
	query := `SELECT * FROM company WHERE inn = $1`

	rows, err := c.Pool.Query(ctx, query, inn)
	if err != nil {
		return entity.Company{}, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	company := entity.Company{}
	for rows.Next() {
		err = rows.Scan(
			&company.ID,
			&company.UUID,
			&company.Name,
			&company.Inn)
		if err != nil {
			return entity.Company{}, fmt.Errorf("error in parsing project: %w", err)
		}
	}
	return company, nil
}
