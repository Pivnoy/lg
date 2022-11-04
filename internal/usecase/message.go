package usecase

import (
	"context"
	"github.com/google/uuid"
	"lg/internal/entity"
)

type MessageUseCase struct {
	repo MessageRp
}

var _ MessageContract = (*MessageUseCase)(nil)

func NewMessageUseCase(repo MessageRp) *MessageUseCase {
	return &MessageUseCase{repo: repo}
}

func (m *MessageUseCase) StoreMessage(ctx context.Context, message entity.Message) error {
	return m.repo.StoreMessage(ctx, message)
}

func (m *MessageUseCase) GetLastMessageByChat(ctx context.Context, chat uuid.UUID) (entity.Message, error) {
	return m.repo.GetLastMessageByChat(ctx, chat)
}
