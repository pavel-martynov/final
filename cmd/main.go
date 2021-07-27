package main

import (
	grpcServer "final/cmd/server"
	"final/config"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.InitConfig()

	if err != nil {
		logrus.Fatal(err)
	}

	grpcServer.StartGRPCServer(*cfg)
}