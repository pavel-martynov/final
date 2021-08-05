package rabbit

import (
	"final/config"
	"fmt"
	"log"

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
	checkErrors(err, fmt.Sprintf("Rabbit dial to %s ", url))

	ch, err := conn.Channel()
	checkErrors(err, "rabbit channel()")

	_, err = ch.QueueDeclare(
		"actions",
		false,
		false,
		false,
		false,
		nil,
	)
	checkErrors(err, "QueueDeclare")

	return conn, nil
}

func checkErrors(err error, details string) {
	if err != nil {
		log.Fatal(details, err)
	}
}
