package server

import (
	"database/sql"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"github.com/shii-cchi/message-processor-go/internal/broker/consumer"
	"github.com/shii-cchi/message-processor-go/internal/broker/producer"
	"github.com/shii-cchi/message-processor-go/internal/config"
	"github.com/shii-cchi/message-processor-go/internal/database"
	"github.com/shii-cchi/message-processor-go/internal/handlers"
	"github.com/shii-cchi/message-processor-go/internal/service"
	"log"
	"fmt"
	"net/http"
)

type Server struct {
	httpServer  *http.Server
	httpHandler *handlers.Handler
}

func NewServer(r chi.Router) (*Server, error) {
	cfg, err := config.LoadConfig()

	if err != nil {
		return nil, err
	}

	conn, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName))

	if err != nil {
		return nil, err
	}

	queries := database.New(conn)

	kafkaProducer, err := producer.NewProducer(cfg.KafkaBroker)

	if err != nil {
		return nil, err
	}

	services := service.NewMessageService(queries, kafkaProducer)

	go func() {
		kafkaConsumer, err := consumer.NewConsumer(cfg.KafkaBroker, services)

		if err != nil {
			log.Fatalf("Failed to create Kafka consumer: %s\n", err)
		}

		kafkaConsumer.StartConsuming()
	}()

	handler := handlers.NewHandler(services)
	handler.RegisterHTTPEndpoints(r)

	log.Printf("Server starting on port %s", cfg.Port)

	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + cfg.Port,
			Handler: r,
		},
		httpHandler: handler,
	}, nil
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}
