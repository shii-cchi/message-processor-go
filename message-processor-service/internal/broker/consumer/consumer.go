package consumer

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/shii-cchi/message-processor-go/internal/service"
	"log"
)

type Consumer struct {
	consumer *kafka.Consumer
	service  *service.MessageService
}

func NewConsumer(kafkaBroker string, s *service.MessageService) (*Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaBroker,
		"group.id":          "processed-messages-group",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to create consumer: %s\n", err)
	}

	err = c.Subscribe("processed-messages", nil)

	if err != nil {
		return nil, fmt.Errorf("Failed to subscribe to topic processed-messages: %s\n", err)
	}

	return &Consumer{
		consumer: c,
		service:  s,
	}, nil
}

func (c Consumer) StartConsuming() {
	go func() {
		for {
			msg, err := c.consumer.ReadMessage(-1)
			if err == nil {
				log.Printf("Received message: %s\n", string(msg.Value))
			} else {
				log.Printf("Consumer error: %v (%v)\n", err, msg)
			}

			err = c.service.UpdateMessageStatus(context.Background(), msg.Value)

			if err != nil {
				log.Println(err)
			}
		}
	}()
}
