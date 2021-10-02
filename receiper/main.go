package main

import (
	"log"
	"rabbit/config"
	"rabbit/constant"
	"rabbit/pkg/message"
	"rabbit/receiper/channel"
)

func main() {
	conn, ch := config.RabbitConfig()
	defer conn.Close()
	defer ch.Close()

	app := message.New().SetRabbitChannel(ch).Validate()

	forever := make(chan bool)
	app.Consume(constant.CatOne, channel.Channel1)
	app.Consume(constant.CatTwo, channel.Channel2)
	app.Consume(constant.CatThree, channel.Channel3)
	app.Consume(constant.CatFour, channel.Channel4)
	app.Consume(constant.CatFive, channel.Channel5)

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<- forever 
}