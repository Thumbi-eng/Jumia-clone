package main

import (
	"context"
	"log"
	"time"

	pb "jumia-clone-backend/services/product-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("Testing Product Service...")
	log.Println("======================")

	// Test 1: Create Product
	log.Println("\n1. Testing Create Product...")
	createResp, err := client.CreateProduct(ctx, &pb.CreateProductRequest{
		Name:               "iPhone 15 Pro",
		Description:        "Latest Apple iPhone with A17 Pro chip",
		Price:              999.99,
		Category:           "Electronics",
		Stock:              50,
		ImageUrl:           "https://example.com/iphone15.jpg",
		Brand:              "Apple",
		DiscountPercentage: 10.0,
	})
	if err != nil {
		log.Fatalf("Create product failed: %v", err)
	}

	log.Printf("✓ Product created! ID: %s", createResp.Id)
	productID := createResp.Id

	// Test 2: Get Product
	log.Println("\n2. Testing Get Product...")
	getResp, err := client.GetProduct(ctx, &pb.GetProductRequest{Id: productID})
	if err != nil {
		log.Fatalf("Get product failed: %v", err)
	}

	log.Printf("✓ Product retrieved!")
	log.Printf("  Name: %s", getResp.Product.Name)
	log.Printf("  Price: $%.2f", getResp.Product.Price)
	log.Printf("  Final Price: $%.2f (after %.0f%% discount)", getResp.Product.FinalPrice, getResp.Product.DiscountPercentage)
	log.Printf("  In Stock: %v (%d units)", getResp.Product.InStock, getResp.Product.Stock)

	// Test 3: List Products
	log.Println("\n3. Testing List Products...")
	listResp, err := client.ListProducts(ctx, &pb.ListProductsRequest{
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		log.Fatalf("List products failed: %v", err)
	}

	log.Printf("✓ Found %d products", listResp.Total)

	// Test 4: Create more products
	log.Println("\n4. Creating more products...")
	products := []struct {
		name     string
		price    float64
		category string
	}{
		{"Samsung Galaxy S24", 899.99, "Electronics"},
		{"Nike Air Max", 129.99, "Shoes"},
		{"Sony Headphones", 249.99, "Electronics"},
	}

	for _, p := range products {
		_, err := client.CreateProduct(ctx, &pb.CreateProductRequest{
			Name:        p.name,
			Description: "Test product",
			Price:       p.price,
			Category:    p.category,
			Stock:       20,
		})
		if err != nil {
			log.Printf("  ✗ Failed to create %s", p.name)
		} else {
			log.Printf("  ✓ Created %s", p.name)
		}
	}

	// Test 5: Search Products
	log.Println("\n5. Testing Search Products...")
	searchResp, err := client.SearchProducts(ctx, &pb.SearchProductsRequest{
		Query:    "phone",
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		log.Fatalf("Search products failed: %v", err)
	}

	log.Printf("✓ Search found %d products matching 'phone'", searchResp.Total)

	// Test 6: Get Products by Category
	log.Println("\n6. Testing Get Products by Category...")
	categoryResp, err := client.GetProductsByCategory(ctx, &pb.GetProductsByCategoryRequest{
		Category: "Electronics",
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		log.Fatalf("Get products by category failed: %v", err)
	}

	log.Printf("✓ Found %d products in Electronics category", categoryResp.Total)

	log.Println("\n======================")
	log.Println("All tests completed! ✓")
}
