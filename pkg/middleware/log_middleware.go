package middleware

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
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

		var reqBody []byte

		contentType := c.GetHeader("Content-Type")
		if c.Request.Body != nil && !strings.HasPrefix(contentType, "multipart/") {
			bodyBytes, _ := ioutil.ReadAll(io.LimitReader(c.Request.Body, 1024))
			reqBody = bodyBytes
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()

		log.Printf("[REQ] %s %s | Body: %s", colorMethod(c.Request.Method), c.Request.URL.Path, truncateForLog(reqBody))
		log.Printf("[RES] %s %s | Status: %s | Duration: %v",
			colorMethod(c.Request.Method), c.Request.URL.Path, colorStatus(status), duration)
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func truncateForLog(b []byte) string {
	if len(b) > 1024 {
		return string(b[:1024]) + "..."
	}
	return string(b)
}
