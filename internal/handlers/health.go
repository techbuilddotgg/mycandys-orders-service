package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// HealthCheck godoc
// @Summary health check
// @Schemes
// @Description do health check
// @Success 200
// @Router /health [get]
// @Tags health
func HealthCheck(g *gin.Context) {

	timestamp := time.Now().Format(time.DateTime)

	health := map[string]string{
		"status":    "ok",
		"timestamp": timestamp,
	}

	g.JSON(http.StatusOK, health)
}
