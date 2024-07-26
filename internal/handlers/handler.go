package handlers

import (
	"github.com/go-chi/chi"
	"github.com/shii-cchi/message-processor-go/internal/service"
)

type Handler struct {
	messageService *service.MessageService
}

func NewHandler(s *service.MessageService) *Handler {
	return &Handler{
		messageService: s,
	}
}

func (h *Handler) RegisterHTTPEndpoints(r chi.Router) {
	r.Post("/messages", h.createMessage)
	r.Get("/messages/stats", h.getStats)
}
