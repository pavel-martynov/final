package router

import (
	"encoding/json"
	"io"
	"net/http"

	"final/internal/message_sender"
	"final/internal/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func PostAction (sender *message_sender.MsgSender) func (w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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

		sender.Send(message)

		w.WriteHeader(http.StatusOK)
	}
}


func NewRouter(sender *message_sender.MsgSender) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/action", PostAction(sender))

	return r
}