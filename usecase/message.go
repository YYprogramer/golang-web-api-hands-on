package usecase

import (
	"context"

	"golang-web-api-hands-on/repository"
)

type Message interface {
	Get(ctx context.Context) string
}

type messageUseCase struct {
	messageRepo repository.MessageRepository
}

func NewMessage(messageRepo repository.MessageRepository) Message {
	return &messageUseCase{
		messageRepo: messageRepo,
	}
}

func (m *messageUseCase) Get(ctx context.Context) string {
	return m.messageRepo.GetMessage()
}
