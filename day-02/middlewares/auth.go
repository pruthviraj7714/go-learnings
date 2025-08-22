package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("authorization")

		if authHeader != "Bearer 123123" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Header not found"})
			c.Abort()
			return
		}

		c.Next()
	}

}
