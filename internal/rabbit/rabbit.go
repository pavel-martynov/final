package rabbit

import (
	"final/config"
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func NewConnection(cfg *config.Config) (*amqp.Connection, error) {
	url := fmt.Sprintf(
		"amqp://%s:%s@%s:%s",
		cfg.Rabbit.User,
		cfg.Rabbit.Password,
		cfg.Rabbit.Addr,
		cfg.Rabbit.Port,
	)
	conn, err := amqp.Dial(url)
	checkErrors(err)

	ch, err := conn.Channel()
	checkErrors(err)
	log.Println("connected ch")

	_, err = ch.QueueDeclare(
		"actions",
		false,
		false,
		false,
		false,
		nil,
	)
	checkErrors(err)

	return conn, nil
}

func checkErrors(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}
