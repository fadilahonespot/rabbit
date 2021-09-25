package config

import (
	"log"

	"github.com/streadway/amqp"
)

func RabbitConfig() (conn *amqp.Connection, ch *amqp.Channel) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
		return 
	}

	ch, err = conn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Failed to open a channel", err)
		return
	}

	return 
}


