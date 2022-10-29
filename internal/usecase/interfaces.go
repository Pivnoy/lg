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
		GetAllProjects(context.Context) ([]entity.Project, error)
		GetProjectByName(context.Context, string) (entity.Project, error)
		CreateProject(context.Context, entity.Project) (string, error)
		UpdateProject(context.Context, entity.Project) error
		DeleteProject(context.Context, string) error
	}

	ProjectContract interface {
		GetAllProjects(context.Context) ([]entity.Project, error)
		GetProjectByName(context.Context, string) (entity.Project, error)
		CreateProject(context.Context, entity.Project) (string, error)
		UpdateProject(context.Context, entity.Project) error
		DeleteProject(context.Context, string) error
	}
)
