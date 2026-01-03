package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"jumia-clone-backend/services/order-service/internal/handler"
	"jumia-clone-backend/services/order-service/internal/models"
	"jumia-clone-backend/services/order-service/internal/repository"
	"jumia-clone-backend/services/order-service/internal/service"
	pb "jumia-clone-backend/services/order-service/proto"
)

func main() {
	// Database connection
	dsn := "host=localhost user=postgres password=postgres dbname=jumia_orders port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate the schema
	if err := db.AutoMigrate(&models.Order{}, &models.OrderItem{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database connected and migrated successfully")

	// Initialize layers
	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	// Set up gRPC server
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, orderHandler)

	log.Println("Order service is running on port 50054...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
