package repo

import (
	"context"
	"fmt"
	"lg/internal/entity"
	"lg/internal/usecase"
	"lg/pkg/postgres"
)

type CategoryRepo struct {
	*postgres.Postgres
}

var _ usecase.CategoryRp = (*CategoryRepo)(nil)

func NewCategoryRepo(pg *postgres.Postgres) *CategoryRepo {
	return &CategoryRepo{pg}
}

func (c *CategoryRepo) GetAllCategory(ctx context.Context) ([]entity.Category, error) {
	query := `SELECT * FROM category`

	rows, err := c.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var categoryList []entity.Category
	for rows.Next() {
		category := entity.Category{}
		err = rows.Scan(
			&category.ID,
			&category.UUID,
			&category.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("error in parsing category: %w", err)
		}
		categoryList = append(categoryList, category)
	}
	return categoryList, nil
}
