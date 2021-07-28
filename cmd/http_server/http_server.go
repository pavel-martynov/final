package httpServer

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"final/config"
	"final/internal/message_sender"
	"final/internal/router"
)

func StartHTTPServer(config config.Config, sender *message_sender.MsgSender) {
	r := router.NewRouter(sender)
	addr := fmt.Sprintf("%s:%s", config.HTTP.Addr, config.HTTP.Port)

	logrus.Info(fmt.Sprintf("Starting HTTP server at %s", addr))

	if err := http.ListenAndServe(addr, r); err != nil {
		logrus.Fatal(err)
	} else {
		fmt.Println("started")
	}
}