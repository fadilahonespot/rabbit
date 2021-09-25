package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"rabbit/constant"
	"rabbit/pkg/message"

	"github.com/gin-gonic/gin"
)

type defaultHandler struct {
	rabbitService message.RabbitService
	router        *gin.Engine
}

func New() *defaultHandler {
	return &defaultHandler{}
}

func (s *defaultHandler) SetRabbitService(t message.RabbitService) *defaultHandler {
	s.rabbitService = t
	return s
}

func (s *defaultHandler) Validate() *defaultHandler {
	if s.rabbitService == nil {
		panic("rabbitService is nil")
	}
	return s
}

func (s *defaultHandler) SendPublish(c *gin.Context) {
	tofic := c.Param("tofic")
	person := constant.Person{}
	err := c.Bind(&person)
	if err != nil {
		log.Panicf("error bind data: %s", err)
	}

	by, _ := json.Marshal(person)
	err = s.rabbitService.Publish(tofic, by)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, "Success Publish")
}
