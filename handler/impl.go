package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"rabbit/constant"
	"rabbit/pkg/message"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
	err := c.ShouldBindBodyWith(&person, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	by, _ := json.Marshal(person)
	err = s.rabbitService.Publish(tofic, by)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "failed publish data")
		return
	}

	c.JSON(http.StatusOK, "Success Publish")
}
