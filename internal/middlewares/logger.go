package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mycandys/orders/internal/env"
	"github.com/mycandys/orders/internal/rabbitmq"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Log struct {
	Timestamp     string `json:"timestamp"`
	CorrelationId string `json:"correlationId"`
	Url           string `json:"url"`
	Message       string `json:"message"`
	Service       string `json:"service"`
	Type          string `json:"type"`
}

func NewLog(correlationId string, url string, message string, service string) *Log {
	return &Log{
		Timestamp:     time.Now().Format(time.RFC3339),
		CorrelationId: correlationId,
		Url:           url,
		Message:       message,
		Service:       service,
	}
}

func (l *Log) setType(logType string) {
	l.Type = logType
}

func (m *Middleware) LogFields(log *Log) *logrus.Entry {
	return m.logger.WithFields(logrus.Fields{
		"timestamp":     log.Timestamp,
		"type":          log.Type,
		"correlationId": log.CorrelationId,
		"url":           log.Url,
		"message":       log.Message,
		"service":       log.Service,
	})
}

func (m *Middleware) LogError(log *Log) {
	m.LogFields(log).Error("Request received")
}

func (m *Middleware) LogWarning(log *Log) {
	m.LogFields(log).Warn("Request received")
}

func (m *Middleware) LogInfo(log *Log) {
	m.LogFields(log).Info("Request received")
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (m *Middleware) Logger() gin.HandlerFunc {
	queueName, _ := env.GetEnvVar(env.QUEUE_NAME)
	exchangeName, _ := env.GetEnvVar(env.EXCHANGE_NAME)

	queue := rabbitmq.DeclareQueue(queueName)

	return func(c *gin.Context) {
		correlationId := c.GetHeader("X-Correlation-Id")

		if correlationId == "" {
			correlationId = uuid.New().String()
		}

		c.Header("X-Correlation-Id", correlationId)

		url := c.Request.Host + c.Request.URL.RequestURI()

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		status := c.Writer.Status()

		method := c.Request.Method

		msg := fmt.Sprintf("%s - %s ", method, c.Request.URL.Path)

		serviceLog := NewLog(correlationId, url, msg, "orders")

		switch {
		case status >= http.StatusInternalServerError:
			serviceLog.setType("error")
			m.LogError(serviceLog)
		case status >= http.StatusBadRequest && status < http.StatusInternalServerError:
			serviceLog.setType("warning")
			m.LogWarning(serviceLog)
		default:
			serviceLog.setType("info")
			m.LogInfo(serviceLog)
		}

		payload, err := json.Marshal(serviceLog)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error marshalling log")
		}

		rabbitmq.Publish(exchangeName, queue.Name, payload, c.Request.Context())
	}
}
