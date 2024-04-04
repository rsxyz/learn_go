package main

import (
	"go-gin-hello/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route for the root endpoint
	router.GET("/", api.HelloHandler)

	// Run the server on port 8080
	router.Run(":8080")
}
