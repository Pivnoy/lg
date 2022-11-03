package usecase

import (
	"context"
	"github.com/google/uuid"
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
		CheckUserExistence(context.Context, string) (bool, error)
	}

	JwtContract interface {
		CompareUserPassword(context.Context, entity.User) error
		GenerateToken(string) (string, error)
		CheckToken(token string) (string, error)
	}

	RegisterContract interface {
		CreateNewUser(context.Context, string, string) (uuid.UUID, error)
	}

	ProjectRp interface {
		GetAllProjects(context.Context) ([]entity.Project, error)
		GetProjectByUUID(context.Context, uuid.UUID) (entity.Project, error)
		CreateProject(context.Context, entity.Project) (uuid.UUID, error)
		DeleteProjectByUUID(context.Context, uuid.UUID) error
	}

	ProjectContract interface {
		GetAllProjects(context.Context) ([]entity.Project, error)
		GetProjectByUUID(context.Context, uuid.UUID) (entity.Project, error)
		CreateProject(context.Context, entity.Project) (uuid.UUID, error)
		DeleteProjectByUUID(context.Context, uuid.UUID) error
		CheckProjectExistenceByProjectUUID(context.Context, uuid.UUID) (bool, error)
	}

	LineupRp interface {
		GetLineupByProjectUUID(context.Context, uuid.UUID) (entity.Lineup, error)
		DeleteLineupByProjectUUID(context.Context, uuid.UUID) error
	}

	LineupContract interface {
		DeleteLineupByProjectUUID(context.Context, uuid.UUID) error
		CheckLineupExistenceByProjectUUID(context.Context, uuid.UUID) (bool, error)
	}

	MessageRp interface {
		StoreMessage(context.Context, entity.Message) (uuid.UUID, error)
	}

	ChatRp interface {
		CreateChat(context.Context)
		AddMessageByChat(context.Context, uuid.UUID, entity.Message) error
		GetCharHistory(context.Context, uuid.UUID) ([]entity.Message, error)
	}

	ProfileRp interface {
		GetProfileByUser(context.Context, uuid.UUID) (entity.Profile, error)
	}

	ProfileContract interface {
		GetProfileByUser(context.Context, uuid.UUID) (entity.Profile, error)
	}

	CountryRp interface {
		GetAllCountries(ctx context.Context) ([]entity.Country, error)
	}

	CountryContract interface {
		GetAllCountries(ctx context.Context) ([]entity.Country, error)
	}
)
