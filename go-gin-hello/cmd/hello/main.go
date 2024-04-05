package main

import (
	"go-gin-hello/internal/api"

	"github.com/gin-gonic/gin"

	"fmt"
	"net/http"
	"os"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route for the root endpoint
	router.GET("/", api.HelloHandler)

	// Check if the application is configured to run with TLS
	if os.Getenv("USE_TLS") == "true" {
		// Ensure TLS_CERT and TLS_KEY environment variables are set
		tlsCert := os.Getenv("TLS_CERT")
		tlsKey := os.Getenv("TLS_KEY")
		if tlsCert == "" || tlsKey == "" {
			fmt.Println("TLS_CERT and TLS_KEY environment variables must be provided for HTTPS requests")
			return
		}

		// Run the server with TLS configuration
		go func() {
			fmt.Println("Listening on port 443...")
			err := http.ListenAndServeTLS(":443", tlsCert, tlsKey, router)
			if err != nil {
				fmt.Printf("Error starting server on port 443: %s\n", err)
			}
		}()
	} else {
		// Run the server on port 8080 without TLS
		fmt.Println("Listening on port 8080...")
		go func() {
			err := http.ListenAndServe(":8080", router)
			if err != nil {
				fmt.Printf("Error starting server on port 8080: %s\n", err)
			}
		}()
	}

	select {}
}
