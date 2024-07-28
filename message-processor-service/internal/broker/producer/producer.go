package producer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

type Producer struct {
	producer *kafka.Producer
}

func NewProducer(kafkaBroker string) (*Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaBroker,
		"acks":              "all"})

	if err != nil {
		return nil, fmt.Errorf("Failed to create producer: %s\n", err)
	}

	go func(p *kafka.Producer) {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("delivery failed: %v\n", ev.TopicPartition.Error)
				} else {
					log.Printf("message delivered to topic %s [%d] at offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			case kafka.Error:
				log.Printf("Kafka error: %v\n", ev)
			default:
				log.Printf("Ignored event: %v\n", ev)
			}
		}
	}(p)

	return &Producer{
		producer: p,
	}, nil
}

func (p *Producer) SendMessage(message []byte) error {
	topic := "unprocessed-messages"

	err := p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, nil)

	if err != nil {
		return fmt.Errorf("Failed to produce message: %s\n", err)
	}

	p.producer.Flush(15 * 1000)

	return nil
}
