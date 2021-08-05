package message_sender

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type MsgSender struct {
	ch        *amqp.Channel
	queueName string
}

func NewMsgSender(ch *amqp.Channel, queueName string) *MsgSender {
	m := MsgSender{
		ch:        ch,
		queueName: queueName,
	}

	return &m
}

func (m *MsgSender) Send(msg interface{}) {
	// for debug you have logrus.Debug
	fmt.Println(msg)
	msgJson, err := json.Marshal(msg)

	// need to return an error
	if err != nil {
		logrus.Error(err)
		return
	}

	// also skipped error, silent mode (please, not hide errors) :)
	m.ch.Publish(
		"",          // exchange
		m.queueName, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/json",
			Body:        msgJson,
		},
	)
}
