package middlewares

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func (m *Middleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		auth := c.GetHeader("Authorization")
		header := strings.Split(auth, " ")

		if len(header) != 2 {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		token := header[1]

		res, err := m.AuthService.ValidateToken(token)
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		c.Set("userId", res.UserId)
		c.Next()
	}
}
