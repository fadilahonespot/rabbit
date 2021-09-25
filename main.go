package main

import (
	"log"

	"rabbit/config"
	"rabbit/handler"
	"rabbit/pkg/message"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "8769"
	route := gin.Default()

	conn, ch := config.RabbitConfig()
	defer conn.Close()
	defer ch.Close()

	rabbitService := message.New().SetRabbitChannel(ch).Validate()
	handler.New().SetRabbitService(rabbitService).SetRouter(route).Validate()

	err := route.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed running app in port: %s : %s", port, err)
	}
}