package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"rabbit/constant"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func SetupMiddleware(server *gin.Engine) {
	server.Use(SetUpLogger())
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "token", "Content-Type", "Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
}

func SetUpLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyLogWriter := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = bodyLogWriter

		//Start time
		startTime := time.Now()
		startNano := startTime.UnixNano() / 1000000

		//Process request
		c.Next()

		var requestBody interface{}
		c.ShouldBindBodyWith(&requestBody, binding.JSON)
		if requestBody == nil {
			requestBody = ""
		}

		//End time
		endTime := time.Now()
		endNano := endTime.UnixNano() / 1000000

		if c.Request.Method == "POST" {
			c.Request.ParseForm()
		}

		//Log format
		accessLogMap := make(map[string]interface{})

		accessLogMap["request_time"] = startTime.Format("2006-01-02 15:04:05")
		accessLogMap["request_method"] = c.Request.Method
		accessLogMap["request_uri"] = c.Request.RequestURI
		accessLogMap["request_ua"] = c.Request.UserAgent()
		accessLogMap["request_body"] = requestBody
		accessLogMap["request_client_host"] = c.Request.Host
		accessLogMap["response_time"] = endTime.Format("2006-01-02 15:04:05")
		accessLogMap["response_code"] = c.Writer.Status()
		accessLogMap["response_data"] = bodyLogWriter.body.String()
		accessLogMap["cost_time"] = fmt.Sprintf("%vms", endNano-startNano)

		accessLogJson, _ := json.Marshal(accessLogMap)

		accessLogName := constant.AccessLogPath + "." + time.Now().Format("20060102")
		if f, err := os.OpenFile(accessLogName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err != nil {
			f, _ := os.Create(accessLogName)
			gin.DefaultWriter = io.MultiWriter(f)
		} else {
			f.WriteString(string(accessLogJson) + "\n")
		}
	}
}

