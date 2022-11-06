package repo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"lg/internal/entity"
	"lg/internal/usecase"
	"lg/pkg/postgres"
)

type MessageRepo struct {
	*postgres.Postgres
}

var _ usecase.MessageRp = (*MessageRepo)(nil)

func NewMessageRepo(pg *postgres.Postgres) *MessageRepo {
	return &MessageRepo{pg}
}

func (m *MessageRepo) StoreMessage(ctx context.Context, message entity.Message) error {
	query := `INSERT INTO message (author_uuid, msg_type, content, creation_date, chat_uuid) VALUES ($1, $2, $3, $4, $5)`

	rows, err := m.Pool.Query(ctx, query, message.AuthorUUID, message.Type, message.Content, message.CreationDate, message.ChatUUID)
	if err != nil {
		return fmt.Errorf("cannot insert value into message: %v", err)
	}
	defer rows.Close()
	return nil
}

func (m *MessageRepo) GetLastMessageByChat(ctx context.Context, chat uuid.UUID) (entity.Message, error) {
	query := `SELECT * FROM message WHERE chat_uuid = $1`

	rows, err := m.Pool.Query(ctx, query, chat)
	if err != nil {
		return entity.Message{}, fmt.Errorf("cannot execute query: %v", err)
	}
	defer rows.Close()
	var msg entity.Message
	for rows.Next() {
		err = rows.Scan(&msg.ID,
			&msg.AuthorUUID,
			&msg.Type,
			&msg.Content,
			&msg.CreationDate,
			&msg.ChatUUID)
		if err != nil {
			return entity.Message{}, fmt.Errorf("cannot parse into message: %v", err)
		}
	}
	return msg, nil
}

func (m *MessageRepo) UpdateMessageStatus(ctx context.Context, user uuid.UUID, chat uuid.UUID) error {
	query := `UPDATE message SET msg_type = 'text' WHERE author_uuid = $1 and chat_uuid = $2`

	rows, err := m.Pool.Query(ctx, query, user, chat)
	if err != nil {
		return fmt.Errorf("cannot execute query: %v", err)
	}
	defer rows.Close()
	return nil
}
