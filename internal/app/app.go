package app

import (
	"github.com/go-chi/chi"
	"github.com/shii-cchi/message-processor-go/internal/server"
	"log"
)

func Run() {
	r := chi.NewRouter()

	s, err := server.NewServer(r)

	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(s.Run())
}
