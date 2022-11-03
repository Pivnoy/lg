package repo

import "lg/pkg/postgres"

type MessageRepo struct {
	*postgres.Postgres
}

func NewMessageRepo(pg *postgres.Postgres) *MessageRepo {
	return &MessageRepo{pg}
}
