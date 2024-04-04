package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// HelloHandler handles the / endpoint
func HelloHandler(c *gin.Context) {
	// Get environment variables MESSAGE and COLOR, use defaults if not set
	message := os.Getenv("MESSAGE")
	if message == "" {
		message = "World"
	}
	color := os.Getenv("COLOR")
	if color == "" {
		color = "black"
	}

	// Construct HTML response with message and color
	htmlResponse := fmt.Sprintf("<h1 style=\"color:%s;\">Hello %s!</h1>", color, message)

	// Send HTML response
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlResponse))
}
