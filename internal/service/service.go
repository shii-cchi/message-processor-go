package service

import (
	"context"
	"github.com/shii-cchi/message-processor-go/internal/database"
)

type MessageService struct {
	Queries *database.Queries
}

func NewMessageService(q *database.Queries) *MessageService {
	return &MessageService{
		Queries: q,
	}
}

func (s MessageService) CreateMessages(ctx context.Context, content string) (database.Message, error) {
	message, err := s.Queries.CreateMessage(ctx, content)

	if err != nil {
		return database.Message{}, err
	}

	return message, nil
}

func (s MessageService) GetStats(ctx context.Context) (int64, int64, error) {
	messagesCount, err := s.Queries.GetMessagesCount(ctx)

	if err != nil {
		return 0, 0, err
	}

	processedMessagesCount, err := s.Queries.GetProcessedMessagesCount(ctx)

	if err != nil {
		return 0, 0, err
	}

	return messagesCount, processedMessagesCount, nil
}
