package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"strings"
	"{{ cookiecutter.project_name }}/parser"
	"{{ cookiecutter.project_name }}/service"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logger(s *service.OperateLogService, log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()

		if strings.ToUpper(c.Request.Method) != "GET" {
			err := s.Create(c, &parser.CreateOperateLogParser{
				Name:       c.GetString("requestName"),
				Path:       c.Request.RequestURI,
				Method:     c.Request.Method,
				StatusCode: c.Writer.Status(),
				IP:         c.ClientIP(),
				UserId:     c.GetInt("userId"),
			})
			if err != nil {
				log.Error(err)
			}
		}
	}
}

func SetRequestNameHandler(requestName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("requestName", requestName)
		c.Next()
	}
}
