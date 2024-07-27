package server

import (
	"context"
	"database/sql"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"github.com/shii-cchi/message-processor-go/internal/broker"
	"github.com/shii-cchi/message-processor-go/internal/config"
	"github.com/shii-cchi/message-processor-go/internal/database"
	"github.com/shii-cchi/message-processor-go/internal/handlers"
	"github.com/shii-cchi/message-processor-go/internal/service"
	"log"
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

	conn, err := sql.Open("postgres", cfg.DbURI)

	if err != nil {
		return nil, err
	}

	queries := database.New(conn)

	kafkaProducer, err := broker.NewProducer()

	if err != nil {
		return nil, err
	}

	kafkaConsumer, err := broker.NewConsumer()

	if err != nil {
		return nil, err
	}

	services := service.NewMessageService(queries, kafkaProducer, kafkaConsumer)

	go func() {
		for {
			msg, err := kafkaConsumer.ReadMessage(-1)
			if err == nil {
				log.Printf("Received message: %s\n", string(msg.Value))
			} else {
				log.Printf("Consumer error: %v (%v)\n", err, msg)
			}

			err = services.UpdateMessageStatus(context.Background(), msg.Value)

			if err != nil {
				log.Println(err)
			}
		}
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
