package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file (if exists)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	listen := os.Getenv("LISTEN")
	port := os.Getenv("PORT")
	trustedProxies := os.Getenv("TRUSTED_PROXIES")

	// Create a Gin router
	r := gin.New()

	// Set trusted proxies from environment variable
	r.SetTrustedProxies([]string{"" + trustedProxies})

	// Apply middlewares
	//__MIDDLEWARES__

	// Routers
	{

		api := r.Group("/api")
		api.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})
	}

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run(listen + ":" + port)
}
