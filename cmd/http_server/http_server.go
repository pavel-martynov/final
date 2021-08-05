package http_server

import (
	"fmt"
	"log"
	"net/http"

	"final/config"
	"final/internal/message_sender"
	"final/internal/router"
)

// for the future, http separete service can be replaced by https://github.com/grpc-ecosystem/grpc-gateway
// It's not easy for understanding now, but keep in mind
func StartHTTPServer(config config.Config, sender *message_sender.MsgSender) {
	r := router.NewRouter(sender)
	addr := fmt.Sprintf("%s:%s", config.HTTP.Addr, config.HTTP.Port)

	log.Println(fmt.Sprintf("Starting HTTP server at %s", addr))

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("started")
	}
}
