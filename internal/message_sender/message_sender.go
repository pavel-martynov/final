package message_sender

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type MsgSender struct {
	ch *amqp.Channel
	queueName string
}

func NewMsgSender (ch *amqp.Channel, queueName string) *MsgSender {
	m := MsgSender{
		ch: ch,
		queueName: queueName,
	}

	return &m
}

func (m *MsgSender) Send (msg interface{}) {
	fmt.Println(msg)
	msgJson, err := json.Marshal(msg)

	if err != nil {
		logrus.Error(err)
		return
	}

	m.ch.Publish(
		"",     // exchange
		m.queueName, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/json",
			Body:        msgJson,
		},
	)
}