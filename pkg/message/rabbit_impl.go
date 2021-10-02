package message

import (
	"bytes"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

type defaultSetting struct {
	ch *amqp.Channel
}

func New() *defaultSetting {
	return &defaultSetting{}
}

func (s *defaultSetting) SetRabbitChannel(ch *amqp.Channel) *defaultSetting {
	s.ch = ch
	return s
}

func (s *defaultSetting) Validate() RabbitService {
	if s.ch == nil {
		panic("rabbit Channel is nil")
	}
	return s
}

func (s *defaultSetting) Publish(topic string, message []byte) (err error) {
	q, err := s.ch.QueueDeclare(
		topic, // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		err = fmt.Errorf("%s : %s", "Failed to declare a queue", err)
		return
	}

	err = s.ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	if err != nil {
		err = fmt.Errorf("%s : %s", "Failed to publish a message", err)
		return
	}
	return nil
}

func (s *defaultSetting) Consume(topic string, consume func(*amqp.Delivery) error) {
	q, err := s.ch.QueueDeclare(
		topic, // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		err = fmt.Errorf("%s : %s", err, "Failed to declare a queue")
		panic(err)
	}

	msgs, err := s.ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		err = fmt.Errorf("%s : %s", err, "Failed to register a consumer")
		panic(err)
	}

	go func() {
		for d := range msgs {
			err = consume(&d)
			if err != nil {
				fmt.Println("error receiper: ", err)
			}
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			d.Ack(false)
		}
	}()
}
