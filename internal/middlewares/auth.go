package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/mycandys/orders/internal/services"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		service := services.NewAuthService()
		token := c.GetHeader("Authorization")

		res, err := service.ValidateToken(token)
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		c.Set("userId", res.UserId)
		c.Next()
	}
}
