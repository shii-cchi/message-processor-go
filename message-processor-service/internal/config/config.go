package config

import (
	"errors"
	"os"
)

type Config struct {
	Port         string
	DbUser string
	DbPassword string
	DbHost string
	DbPort string
	DbName string
	KafkaBroker string
}

func LoadConfig() (*Config, error) {
	port := os.Getenv("PORT")

	if port == "" {
		return nil, errors.New("PORT is not found")
	}

	dbUser := os.Getenv("DB_USER")

	if dbUser == "" {
		return nil, errors.New("DB_USER is not found")
	}

	dbPassword := os.Getenv("DB_PASSWORD")

	if dbPassword == "" {
		return nil, errors.New("DB_PASSWORD is not found")
	}

	dbHost := os.Getenv("DB_HOST")

	if dbHost == "" {
		return nil, errors.New("DB_HOST is not found")
	}

    dbPort := os.Getenv("DB_PORT")

	if dbPort == "" {
		return nil, errors.New("DB_PORT is not found")
	}

    dbName := os.Getenv("DB_NAME")

	if dbName == "" {
		return nil, errors.New("DB_NAME is not found")
	}

	kafkaBroker := os.Getenv("KAFKA_BROKER")

	if kafkaBroker == "" {
		return nil, errors.New("KAFKA_BROKER is not found")
	}

	return &Config{
		Port:         port,
		DbUser: dbUser,
		DbPassword: dbPassword,
		DbHost: dbHost,
		DbPort: dbPort,
		DbName: dbName,
		KafkaBroker: kafkaBroker,
	}, nil
}
