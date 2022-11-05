package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"lg/internal/entity"
)

type ChatUseCase struct {
	repo      ChatRp
	mContract MessageContract
}

var _ ChatContract = (*ChatUseCase)(nil)

func NewChatUseCase(repo ChatRp, mContract MessageContract) *ChatUseCase {
	return &ChatUseCase{repo: repo, mContract: mContract}
}

func (c *ChatUseCase) CreateChat(ctx context.Context, chatName string, userUUIDs []uuid.UUID) (uuid.UUID, error) {
	chatUUID := uuid.New()
	err := c.repo.CreateChat(ctx, entity.Chat{UUID: chatUUID, Name: chatName})
	if err != nil {
		return uuid.Nil, err
	}
	for _, userUUID := range userUUIDs {
		err := c.repo.AddUserIntoChat(ctx, userUUID, chatUUID)
		if err != nil {
			return uuid.Nil, fmt.Errorf("cannot add user into chat: %v", err)
		}
	}
	return chatUUID, nil
}

func (c *ChatUseCase) GetAllChatsByUser(ctx context.Context, user uuid.UUID) ([]entity.ChatItem, error) {
	chats, err := c.repo.GetAllChatsByUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("eror in getting chat list: %v", err)
	}
	var lastMessages []entity.Message
	for _, chat := range chats {
		msg, err := c.mContract.GetLastMessageByChat(ctx, chat.UUID)
		if err != nil {
			return nil, err
		}
		lastMessages = append(lastMessages, msg)
	}
	var chatItems []entity.ChatItem
	for j := 0; j < len(chats); j++ {
		chatItems = append(chatItems, entity.ChatItem{
			ChatName:    chats[j].Name,
			ChatUUID:    chats[j].UUID,
			ProjectUUID: chats[j].ProjectUUID,
			LastMessage: lastMessages[j],
		})
	}
	return chatItems, nil
}

func (c *ChatUseCase) GetChatHistory(ctx context.Context, chat uuid.UUID) ([]entity.Message, error) {
	return c.repo.GetChatHistory(ctx, chat)
}
