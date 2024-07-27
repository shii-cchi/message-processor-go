package broker

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewConsumer() (*kafka.Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
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

	return c, nil
}
