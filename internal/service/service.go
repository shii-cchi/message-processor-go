package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
	"github.com/shii-cchi/message-processor-go/internal/broker"
	"github.com/shii-cchi/message-processor-go/internal/database"
)

type MessageService struct {
	queries       *database.Queries
	kafkaProducer *kafka.Producer
	kafkaConsumer *kafka.Consumer
}

type Message struct {
	ID      uuid.UUID `json:"id"`
	Content string    `json:"content"`
}

func NewMessageService(q *database.Queries, p *kafka.Producer, c *kafka.Consumer) *MessageService {
	return &MessageService{
		queries:       q,
		kafkaProducer: p,
		kafkaConsumer: c,
	}
}

func (s MessageService) CreateMessages(ctx context.Context, content string) (database.Message, error) {
	message, err := s.queries.CreateMessage(ctx, database.CreateMessageParams{
		ID:      uuid.New(),
		Content: content,
	})

	if err != nil {
		return database.Message{}, err
	}

	return message, nil
}

func (s MessageService) SendMessageToKafka(message database.Message) error {
	msg, err := json.Marshal(Message{
		ID:      message.ID,
		Content: message.Content,
	})

	if err != nil {
		return fmt.Errorf("failed to serialize message: %s", err)
	}

	err = broker.SendMessage(s.kafkaProducer, msg)

	return err
}

func (s MessageService) UpdateMessageStatus(ctx context.Context, message []byte) error {
	var msg Message

	if err := json.Unmarshal(message, &msg); err != nil {
		return fmt.Errorf("failed to unmarshal message: %s", err)
	}

	err := s.queries.UpdateMessageStatus(ctx, database.UpdateMessageStatusParams{
		ID:     msg.ID,
		Status: "processed",
	})

	if err != nil {
		return fmt.Errorf("failed to update message status: %s", err)
	}

	return nil
}

func (s MessageService) GetStats(ctx context.Context) (int64, int64, error) {
	messagesCount, err := s.queries.GetMessagesCount(ctx)

	if err != nil {
		return 0, 0, err
	}

	processedMessagesCount, err := s.queries.GetProcessedMessagesCount(ctx)

	if err != nil {
		return 0, 0, err
	}

	return messagesCount, processedMessagesCount, nil
}
