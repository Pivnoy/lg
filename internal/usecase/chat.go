package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"lg/internal/entity"
)

type ChatUseCase struct {
	repo ChatRp
}

var _ ChatContract = (*ChatUseCase)(nil)

func NewChatUseCase(repo ChatRp) *ChatUseCase {
	return &ChatUseCase{repo: repo}
}

func (c *ChatUseCase) CreateChat(ctx context.Context, chatName string, userUUIDs []uuid.UUID) error {
	chatUUID := uuid.New()
	err := c.repo.CreateChat(ctx, entity.Chat{UUID: chatUUID, Name: chatName})
	if err != nil {
		return err
	}
	for _, userUUID := range userUUIDs {
		err := c.repo.AddUserIntoChat(ctx, userUUID, chatUUID)
		if err != nil {
			return fmt.Errorf("cannot add user into chat: %v", err)
		}
	}
	return nil
}
