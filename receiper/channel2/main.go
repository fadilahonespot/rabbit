package main

import (
	"fmt"
	"rabbit/constant"
	"log"
	"rabbit/config"
	"rabbit/pkg/message"
)

func main() {
	conn, ch := config.RabbitConfig()
	defer conn.Close()
	defer ch.Close()

	rabbitService := message.New().SetRabbitChannel(ch).Validate()

	err := rabbitService.Consume(constant.CatTwo, func(i []byte) error {
		if i == nil {
			fmt.Println("Error")
			return fmt.Errorf("value is nil")
		}
		fmt.Println("THIS CHANNEL 2: ", string(i))
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}