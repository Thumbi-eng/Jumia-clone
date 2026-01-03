package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"jumia-clone-backend/services/cart-service/internal/handler"
	"jumia-clone-backend/services/cart-service/internal/models"
	"jumia-clone-backend/services/cart-service/internal/repository"
	"jumia-clone-backend/services/cart-service/internal/service"
	pb "jumia-clone-backend/services/cart-service/proto"
)

func main() {
	// Database connection
	dsn := "host=localhost user=postgres password=postgres dbname=jumia_carts port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate the schema
	if err := db.AutoMigrate(&models.Cart{}, &models.CartItem{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database connected and migrated successfully")

	// Initialize layers
	cartRepo := repository.NewCartRepository(db)
	cartService := service.NewCartService(cartRepo)
	cartHandler := handler.NewCartHandler(cartService)

	// Set up gRPC server
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCartServiceServer(s, cartHandler)

	log.Println("Cart service is running on port 50053...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
