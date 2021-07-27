package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type RabbitConfig struct {
	Addr string `json:"address"`
	Port string `json:"port"`
}

type GRPCConfig struct {
	Addr string `json:"address"`
	Port string `json:"port"`
}

type Config struct {
	Debug		bool		`json:"debug"`
	Rabbit		RabbitConfig 	`json:"rabbit"`
	GRPC	GRPCConfig 	`json:"grpc"`
}

func InitConfig() (*Config, error) {
	f, err := os.Open("./config.json")

	if err != nil {
		fmt.Println(os.Getwd())
		log.Fatal(err)

		return nil, err
	}

	data, err := io.ReadAll(f)

	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	c := Config{}

	err = json.Unmarshal(data, &c)

	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	return &c, nil
}