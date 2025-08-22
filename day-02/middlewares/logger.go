package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		log.Printf("%s %s %d %d", c.Request.Method, c.Request.URL, c.Writer.Status(), time.Since(start))
	}
}
