package message

import "github.com/streadway/amqp"

type RabbitService interface {
	Publish(topic string, message []byte) (err error)
	Consume(topic string, consume func(*amqp.Delivery) error)
}
