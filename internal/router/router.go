package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"final/internal/http_handlers"
	"final/internal/message_sender"
)

func NewRouter(sender *message_sender.MsgSender) http.Handler {
	r := chi.NewRouter()
	h := http_handlers.NewHandlers(sender)
	r.Use(middleware.Logger)
	r.Post("/action", h.PostAction)

	return r
}
