package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	grpc_server "final/cmd/grpc_server"
	http_server "final/cmd/http_server"
	"final/config"
	"final/internal/message_sender"
	"final/internal/rabbit"
)

func setupCloseHandler(callback func()) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("\nCtrl-C pressed in terminal. Aborting...")
		fmt.Println("ENDING")
		callback()
	}()
}

func main() {
	wg := sync.WaitGroup{}
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

	wg.Add(1)
	setupCloseHandler(wg.Done)
	wg.Wait()
}
