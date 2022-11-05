package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"lg/internal/entity"
)

type CompanyUseCase struct {
	repo CompanyRp
}

var _ CompanyContract = (*CompanyUseCase)(nil)

func NewCompanyUseCase(repo CompanyRp) *CompanyUseCase {
	return &CompanyUseCase{repo: repo}
}

func (c *CompanyUseCase) CreateCompany(ctx context.Context, company entity.Company) (uuid.UUID, error) {
	return c.CreateCompany(ctx, company)
}

func (c *CompanyUseCase) CheckCompanyExistenceByInn(ctx context.Context, inn string) (bool, error) {
	company, err := c.GetCompanyByInn(ctx, inn)
	if err != nil {
		if (err.Error() == "there is no company with this inn") && (company == entity.Company{}) {
			return true, nil
		}
		return false, err
	}
	return false, nil

}

func (c *CompanyUseCase) GetCompanyByInn(ctx context.Context, inn string) (entity.Company, error) {
	company, err := c.repo.GetCompanyByInn(ctx, inn)
	if (err == nil) && (company == entity.Company{}) {
		return company, fmt.Errorf("there is no company with this inn")
	}
	return company, err
}
