package repo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"lg/internal/entity"
	"lg/internal/usecase"
	"lg/pkg/postgres"
)

type ChatRepo struct {
	*postgres.Postgres
}

var _ usecase.ChatRp = (*ChatRepo)(nil)

func NewChatRepo(pg *postgres.Postgres) *ChatRepo {
	return &ChatRepo{pg}
}

func (c *ChatRepo) CreateChat(ctx context.Context, chat entity.Chat) error {
	query := `INSERT INTO chat(uuid, name) VALUES ($1, $2)`

	rows, err := c.Pool.Query(ctx, query, chat.UUID, chat.Name)
	if err != nil {
		return fmt.Errorf("cannot insert values into chat: %v", err)
	}
	defer rows.Close()
	return nil
}

func (c *ChatRepo) AddUserIntoChat(ctx context.Context, user uuid.UUID, chat uuid.UUID) error {
	query := `INSERT INTO chat_member(user_uuid, chat_uuid) VALUES ($1, $2)`

	rows, err := c.Pool.Query(ctx, query, user, chat)
	if err != nil {
		return fmt.Errorf("cannot insert value into chat_member: %v", err)
	}
	defer rows.Close()
	return nil
}

func (c *ChatRepo) GetChatHistory(ctx context.Context, chat uuid.UUID) ([]entity.Message, error) {
	query := `SELECT * FROM message WHERE chat_uuid = $1`

	rows, err := c.Pool.Query(ctx, query, chat)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %v", err)
	}
	defer rows.Close()

	var messages []entity.Message
	for rows.Next() {
		var message entity.Message
		err = rows.Scan(&message.ID,
			&message.AuthorUUID,
			&message.Content,
			&message.CreationDate,
			&message.ChatUUID)
		if err != nil {
			return nil, fmt.Errorf("cannot parse message: %v", err)
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func (c *ChatRepo) GetAllChatsByUser(ctx context.Context, user uuid.UUID) ([]entity.Chat, error) {
	query := `select chat.id, chat.uuid, chat.name, chat.project_uuid from chat
    join chat_member on chat.uuid = chat_member.chat_uuid
    where chat_member.user_uuid = $1`

	rows, err := c.Pool.Query(ctx, query, user)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %v", err)
	}
	defer rows.Close()

	var chats []entity.Chat
	for rows.Next() {
		var chat entity.Chat
		err = rows.Scan(&chat.ID,
			&chat.UUID,
			&chat.Name,
			&chat.ProjectUUID)
		if err != nil {
			return nil, fmt.Errorf("cannot parse chat entity: %v", err)
		}
		chats = append(chats, chat)
	}
	return chats, nil
}
