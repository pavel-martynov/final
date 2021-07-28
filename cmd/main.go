package main

import (
	httpServer "final/cmd/http_server"
	"final/internal/message_sender"
	"final/internal/rabbit"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	grpcServer "final/cmd/grps_server"
	"final/config"
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
		logrus.Fatal(err)
	}

	conn, err := rabbit.NewConn()


	if err != nil {
		logrus.Fatal(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		logrus.Fatal(err)
	}

	defer ch.Close()

	msgSender := message_sender.NewMsgSender(ch, "actions")

	go grpcServer.StartGRPCServer(*cfg, msgSender)

	go httpServer.StartHTTPServer(*cfg, msgSender)

	wg.Add(1)
	setupCloseHandler(wg.Done)
	wg.Wait()
}
