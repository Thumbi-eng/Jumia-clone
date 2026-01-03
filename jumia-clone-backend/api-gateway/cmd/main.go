package main

import (
	"log"
	"os"

	"jumia-clone-backend/api-gateway/internal/client"
	"jumia-clone-backend/api-gateway/internal/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Get configuration from environment
	userServiceAddr := getEnv("USER_SERVICE_ADDR", "localhost:50051")
	port := getEnv("PORT", "8080")

	// Connect to user service
	userClient, err := client.NewUserServiceClient(userServiceAddr)
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}
	defer userClient.Close()

	// Create Gin router
	router := gin.Default()

	// Configure CORS for frontend
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Register routes
	handler.RegisterRoutes(router, userClient.Conn)

	// Start HTTP server
	log.Printf("API Gateway starting on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// getEnv gets environment variable or returns default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
