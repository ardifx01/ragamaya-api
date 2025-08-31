package middleware

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
)

func colorMethod(method string) string {
	switch method {
	case "GET":
		return Blue + method + Reset
	case "POST":
		return Green + method + Reset
	case "PUT":
		return Yellow + method + Reset
	case "DELETE":
		return Red + method + Reset
	default:
		return Cyan + method + Reset
	}
}

func colorStatus(status int) string {
	switch {
	case status >= 200 && status < 300:
		return Green + fmt.Sprintf("%d", status) + Reset
	case status >= 400 && status < 500:
		return Yellow + fmt.Sprintf("%d", status) + Reset
	case status >= 500:
		return Red + fmt.Sprintf("%d", status) + Reset
	default:
		return fmt.Sprintf("%d", status)
	}
}

func RequestResponseLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()

		log.Printf("[REQ] %s %s", colorMethod(c.Request.Method), c.Request.URL.Path)
		log.Printf("[RES] %s %s | Status: %s | Duration: %v",
			colorMethod(c.Request.Method), c.Request.URL.Path, colorStatus(status), duration)
	}
}