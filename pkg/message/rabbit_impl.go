package message

import (
	"bytes"
	"fmt"
	"log"
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

func (s *defaultSetting) Publish(topick string, message []byte) (err error) {
	q, err := s.ch.QueueDeclare(
		topick, // name
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

func (s *defaultSetting) Consume(topick string, f func([]byte) error) (err error) {
	q, err := s.ch.QueueDeclare(
		topick, // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		err = fmt.Errorf("%s : %s", err, "Failed to declare a queue")
		return
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
		return
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			f(d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
	return
}
