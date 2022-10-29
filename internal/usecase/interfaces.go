package usecase

import (
	"context"
	"lg/internal/entity"
)

type (
	UserRp interface {
		GetUserByEmail(context.Context, string) (entity.User, error)
		StoreUser(context.Context, entity.User) error
	}

	UserContract interface {
		GetUser(context.Context, string) (entity.User, error)
		StoreUser(context.Context, entity.User) error
	}

	ProjectRp interface {
		GetProjectByName(context.Context, string) (entity.Project, error)
	}

	ProjectContract interface {
		GetProjectByName(context.Context, string) (entity.Project, error)
	}
)
