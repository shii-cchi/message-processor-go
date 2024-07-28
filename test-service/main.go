package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

func main() {
	log.Println("test-service starting")

	kafkaBroker := os.Getenv("KAFKA_BROKER")

	if kafkaBroker == "" {
		log.Fatal("KAFKA_BROKER is not found")
	}

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaBroker,
		"group.id":          "unprocessed-messages-group",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatalf("Failed to create consumer: %s\n", err)
	}

	defer c.Close()

	err = c.Subscribe("unprocessed-messages", nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s\n", err)
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaBroker})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	defer p.Close()

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					log.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	msgChan := make(chan *kafka.Message)

	go func() {
		for {
			msg, err := c.ReadMessage(-1)
			if err == nil {
				log.Printf("Received message: %s\n", string(msg.Value))
			} else {
				log.Printf("Consumer error: %v (%v)\n", err, msg)
			}

			msgChan <- msg
		}
	}()

	topic := "processed-messages"
	for msg := range msgChan {
		err := p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          msg.Value,
		}, nil)

		if err != nil {
			log.Printf("Failed to produce message: %s\n", err)
		}
	}

	p.Flush(15 * 1000)
}
