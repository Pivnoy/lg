package repo

import "lg/pkg/postgres"

type ChatRepo struct {
	*postgres.Postgres
}

func NewChatRepo(pg *postgres.Postgres) *ChatRepo {
	return &ChatRepo{pg}
}
