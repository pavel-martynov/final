package rabbit


import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func NewConn() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	checkErrors(err)

	ch, err := conn.Channel()
	checkErrors(err)
	fmt.Println("connected ch")

	_, err = ch.QueueDeclare(
		"actions",
		false,
		false,
		false,
		false,
		nil,
	)
	checkErrors(err)

	/*q2, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)

	body1 := "[mes1] Hello, Students!"
	body2 := "[mes2] Hello, Students!"

	err = ch.Publish(
		"", q1.Name, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body1),
		},
	)

	err = ch.Publish(
		"", q2.Name, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body2),
		},
	)

	checkErrors(err)
	*/

	return conn, nil
}

func checkErrors(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}