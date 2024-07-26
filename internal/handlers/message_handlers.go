package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/shii-cchi/message-processor-go/internal/handlers/dto"
	"net/http"
)

func (h *Handler) createMessage(w http.ResponseWriter, r *http.Request) {
	newMessage := dto.MessageDto{}

	err := json.NewDecoder(r.Body).Decode(&newMessage)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	message, err := h.messageService.CreateMessages(r.Context(), newMessage.Content)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Coudn't create message: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, message)
}

func (h *Handler) getStats(w http.ResponseWriter, r *http.Request) {
	messagesCount, processedMessagesCount, err := h.messageService.GetStats(r.Context())

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Coudn't get stats: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, dto.StatsDto{
		MessagesCount:          messagesCount,
		ProcessedMessagesCount: processedMessagesCount,
	})
}
