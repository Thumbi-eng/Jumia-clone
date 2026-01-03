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
	productServiceAddr := getEnv("PRODUCT_SERVICE_ADDR", "localhost:50052")
	cartServiceAddr := getEnv("CART_SERVICE_ADDR", "localhost:50053")
	orderServiceAddr := getEnv("ORDER_SERVICE_ADDR", "localhost:50054")
	port := getEnv("PORT", "8080")

	// Connect to user service
	userClient, err := client.NewUserServiceClient(userServiceAddr)
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}
	defer userClient.Close()

	// Connect to product service
	productClient, err := client.NewProductServiceClient(productServiceAddr)
	if err != nil {
		log.Fatalf("Failed to connect to product service: %v", err)
	}
	defer productClient.Close()

	// Connect to cart service
	cartClient, err := client.NewCartServiceClient(cartServiceAddr)
	if err != nil {
		log.Fatalf("Failed to connect to cart service: %v", err)
	}
	defer cartClient.Close()

	// Connect to order service
	orderClient, err := client.NewOrderServiceClient(orderServiceAddr)
	if err != nil {
		log.Fatalf("Failed to connect to order service: %v", err)
	}
	defer orderClient.Close()

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

	// Register routes for all services
	handler.RegisterAllRoutes(router, userClient.Conn, productClient.Conn, cartClient.Conn, orderClient.Conn)

	// Start HTTP server
	log.Printf("API Gateway starting on port %s...", port)
	log.Printf("Connected to services:")
	log.Printf("  - User Service: %s", userServiceAddr)
	log.Printf("  - Product Service: %s", productServiceAddr)
	log.Printf("  - Cart Service: %s", cartServiceAddr)
	log.Printf("  - Order Service: %s", orderServiceAddr)

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
