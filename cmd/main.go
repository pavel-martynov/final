package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"final/cmd/grpc_server"
	"final/cmd/http_server"
	"final/config"
	"final/internal/message_sender"
	"final/internal/rabbit"
)

func setupTerminationHandler() chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	return c
}

func main() {
	c := setupTerminationHandler()
	cfg, err := config.InitConfig()

	if err != nil {
		log.Fatal("Init config", err)
	}

	conn, err := rabbit.NewConnection(cfg)

	if err != nil {
		log.Fatal("Rabbit new connection", err)
	}

	defer conn.Close()

	rabbitCh, err := conn.Channel()
	if err != nil {
		log.Fatal("Rabbit channel", err)
	}

	defer rabbitCh.Close()

	msgSender := message_sender.NewMsgSender(rabbitCh, "actions")

	go grpc_server.StartGRPCServer(*cfg, msgSender)

	go http_server.StartHTTPServer(*cfg, msgSender)

	<-c
}
