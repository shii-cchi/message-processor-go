package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port         string
	DbURI        string
	KafkaBrokers string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return nil, err
	}

	port := os.Getenv("PORT")

	if port == "" {
		return nil, errors.New("PORT is not found")
	}

	dbURI := os.Getenv("DB_URI")

	if dbURI == "" {
		return nil, errors.New("DB_URL is not found")
	}

	kafkaBrokers := os.Getenv("KAFKA_BROKERS")

	if kafkaBrokers == "" {
		return nil, errors.New("KAFKA_BROKERS is not found")
	}

	return &Config{
		Port:         port,
		DbURI:        dbURI,
		KafkaBrokers: kafkaBrokers,
	}, nil
}
