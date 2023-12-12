package middlewares

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (m *Middleware) Logger() gin.HandlerFunc {

	return func(c *gin.Context) {
		correlationId := c.GetHeader("X-Correlation-Id")

		if correlationId == "" {
			correlationId = uuid.New().String()
		}

		c.Header("X-Correlation-Id", correlationId)

		timestamp := time.Now().Format(time.RFC3339)

		fullUrl := c.Request.Host + c.Request.URL.RequestURI()

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		status := c.Writer.Status()
		switch {
		case status >= http.StatusInternalServerError:
			m.logger.WithFields(logrus.Fields{
				"timestamp":     timestamp,
				"logType":       "Error",
				"url":           fullUrl,
				"correlationId": correlationId,
				"service":       "orders",
			}).Error("Request received")
		case status >= http.StatusBadRequest && status < http.StatusInternalServerError:
			m.logger.WithFields(logrus.Fields{
				"timestamp":     timestamp,
				"logType":       "Warning",
				"url":           fullUrl,
				"correlationId": correlationId,
				"service":       "orders",
			}).Warning("Request received")
		default:
			m.logger.WithFields(logrus.Fields{
				"timestamp":     timestamp,
				"logType":       "Info",
				"url":           fullUrl,
				"correlationId": correlationId,
				"service":       "orders",
			}).Info("Request received")
		}

	}
}
