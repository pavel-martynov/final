package http_handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"final/internal/message_sender"
	"final/internal/model"
)

type Handlers struct {
	sender *message_sender.MsgSender
}

func (h *Handlers) PostAction(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var message model.Action

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.Unmarshal(b, &message); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	h.sender.Send(message)

	w.WriteHeader(http.StatusOK)
}

func NewHandlers(sender *message_sender.MsgSender) *Handlers {
	return &Handlers{sender: sender}
}
