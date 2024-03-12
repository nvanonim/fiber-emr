package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequestLogger is a middleware function that logs the details of each incoming request.
// if body contains password, delete it
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read request body
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// Check if request body contains "password"
		if bytes.Contains(bodyBytes, []byte("password")) {
			var bodyMap map[string]interface{}
			if err := json.Unmarshal(bodyBytes, &bodyMap); err == nil {
				// Redact password if found
				if _, ok := bodyMap["password"]; ok {
					bodyMap["password"] = "********"
				}
				// Marshal modified body back to bytes
				bodyBytes, _ = json.Marshal(bodyMap)
			}
		}

		// Log request details
		log.Println("==================================================")
		log.Printf("Request: %s %s", c.Request.Method, c.Request.RequestURI)
		log.Printf("Header: %v", c.Request.Header)
		log.Printf("Body: %s", string(bodyBytes))

		// Call the next handler
		c.Next()
	}
}

// ResponseLogger is a middleware function that logs the details of each outgoing response.
func ResponseLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a response writer to capture the response
		w := &responseLogger{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}

		// Replace the gin context's writer with the custom response writer
		c.Writer = w

		// Call the next handler
		c.Next()

		// Log response details
		log.Printf("Response: %d with Body: %s", w.status, w.body.String())
	}
}

// Custom response writer to capture the response body
type responseLogger struct {
	gin.ResponseWriter
	body   *bytes.Buffer
	status int
}

// WriteHeader captures the status code of the response
func (w *responseLogger) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

// Write captures the response body
func (w *responseLogger) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = http.StatusOK
	}
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
