package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LatencyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
