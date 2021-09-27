package handler

import (
	"github.com/gin-gonic/gin"
)

func (s *defaultHandler) SetRouter(t *gin.Engine) *defaultHandler {
	SetupMiddleware(t)
	
	t.POST("/publish/:tofic", s.SendPublish)

	s.router = t
	return s
}
