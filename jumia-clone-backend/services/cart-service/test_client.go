package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "jumia-clone-backend/services/cart-service/proto"
)

const (
	address        = "localhost:50053"
	testUserID     = "550e8400-e29b-41d4-a716-446655440000"
	testProductID1 = "660e8400-e29b-41d4-a716-446655440001"
	testProductID2 = "660e8400-e29b-41d4-a716-446655440002"
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCartServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	log.Println("===== Cart Service Client Test =====\n")

	// Test 1: Add first product to cart
	log.Println("Test 1: Adding first product to cart...")
	addResp1, err := client.AddToCart(ctx, &pb.AddToCartRequest{
		UserId:      testUserID,
		ProductId:   testProductID1,
		Quantity:    2,
		Price:       29.99,
		ProductName: "Smartphone",
		ImageUrl:    "https://example.com/smartphone.jpg",
	})
	if err != nil {
		log.Fatalf("AddToCart failed: %v", err)
	}
	log.Printf("Success: %v, Message: %s\n", addResp1.Success, addResp1.Message)
	log.Printf("Cart has %d items, Total: $%.2f\n\n", addResp1.Cart.TotalItems, addResp1.Cart.TotalPrice)

	// Test 2: Add second product to cart
	log.Println("Test 2: Adding second product to cart...")
	addResp2, err := client.AddToCart(ctx, &pb.AddToCartRequest{
		UserId:      testUserID,
		ProductId:   testProductID2,
		Quantity:    1,
		Price:       15.50,
		ProductName: "Headphones",
		ImageUrl:    "https://example.com/headphones.jpg",
	})
	if err != nil {
		log.Fatalf("AddToCart failed: %v", err)
	}
	log.Printf("Success: %v, Message: %s\n", addResp2.Success, addResp2.Message)
	log.Printf("Cart has %d items, Total: $%.2f\n\n", addResp2.Cart.TotalItems, addResp2.Cart.TotalPrice)

	// Test 3: Get cart
	log.Println("Test 3: Getting cart...")
	getResp, err := client.GetCart(ctx, &pb.GetCartRequest{
		UserId: testUserID,
	})
	if err != nil {
		log.Fatalf("GetCart failed: %v", err)
	}
	log.Printf("Success: %v, Message: %s\n", getResp.Success, getResp.Message)
	log.Printf("Cart ID: %s, Total Items: %d, Total Price: $%.2f\n", getResp.Cart.Id, getResp.Cart.TotalItems, getResp.Cart.TotalPrice)
	log.Println("Items:")
	for i, item := range getResp.Cart.Items {
		log.Printf("  %d. %s (x%d) - $%.2f each, Subtotal: $%.2f\n", i+1, item.ProductName, item.Quantity, item.Price, item.Subtotal)
	}
	log.Println()

	// Test 4: Update cart item quantity
	log.Println("Test 4: Updating first product quantity to 5...")
	updateResp, err := client.UpdateCartItem(ctx, &pb.UpdateCartItemRequest{
		UserId:    testUserID,
		ProductId: testProductID1,
		Quantity:  5,
	})
	if err != nil {
		log.Fatalf("UpdateCartItem failed: %v", err)
	}
	log.Printf("Success: %v, Message: %s\n", updateResp.Success, updateResp.Message)
	log.Printf("Cart has %d items, Total: $%.2f\n\n", updateResp.Cart.TotalItems, updateResp.Cart.TotalPrice)

	// Test 5: Add same product again (should merge)
	log.Println("Test 5: Adding first product again (quantity +3)...")
	addResp3, err := client.AddToCart(ctx, &pb.AddToCartRequest{
		UserId:      testUserID,
		ProductId:   testProductID1,
		Quantity:    3,
		Price:       29.99,
		ProductName: "Smartphone",
		ImageUrl:    "https://example.com/smartphone.jpg",
	})
	if err != nil {
		log.Fatalf("AddToCart failed: %v", err)
	}
	log.Printf("Success: %v, Message: %s\n", addResp3.Success, addResp3.Message)
	log.Printf("Cart has %d items (should be 9 total), Total: $%.2f\n\n", addResp3.Cart.TotalItems, addResp3.Cart.TotalPrice)

	// Test 6: Remove second product from cart
	log.Println("Test 6: Removing headphones from cart...")
	removeResp, err := client.RemoveFromCart(ctx, &pb.RemoveFromCartRequest{
		UserId:    testUserID,
		ProductId: testProductID2,
	})
	if err != nil {
		log.Fatalf("RemoveFromCart failed: %v", err)
	}
	log.Printf("Success: %v, Message: %s\n", removeResp.Success, removeResp.Message)
	log.Printf("Cart has %d items, Total: $%.2f\n\n", removeResp.Cart.TotalItems, removeResp.Cart.TotalPrice)

	// Test 7: Get final cart state
	log.Println("Test 7: Getting final cart state...")
	getFinalResp, err := client.GetCart(ctx, &pb.GetCartRequest{
		UserId: testUserID,
	})
	if err != nil {
		log.Fatalf("GetCart failed: %v", err)
	}
	log.Printf("Cart ID: %s, Total Items: %d, Total Price: $%.2f\n", getFinalResp.Cart.Id, getFinalResp.Cart.TotalItems, getFinalResp.Cart.TotalPrice)
	log.Println("Final Items:")
	for i, item := range getFinalResp.Cart.Items {
		log.Printf("  %d. %s (x%d) - $%.2f each, Subtotal: $%.2f\n", i+1, item.ProductName, item.Quantity, item.Price, item.Subtotal)
	}
	log.Println()

	// Test 8: Clear cart
	log.Println("Test 8: Clearing cart...")
	clearResp, err := client.ClearCart(ctx, &pb.ClearCartRequest{
		UserId: testUserID,
	})
	if err != nil {
		log.Fatalf("ClearCart failed: %v", err)
	}
	log.Printf("Success: %v, Message: %s\n\n", clearResp.Success, clearResp.Message)

	// Test 9: Verify cart is empty
	log.Println("Test 9: Verifying cart is empty...")
	getEmptyResp, err := client.GetCart(ctx, &pb.GetCartRequest{
		UserId: testUserID,
	})
	if err != nil {
		log.Fatalf("GetCart failed: %v", err)
	}
	log.Printf("Cart has %d items, Total: $%.2f\n", getEmptyResp.Cart.TotalItems, getEmptyResp.Cart.TotalPrice)

	log.Println("\n===== All Tests Completed Successfully! =====")
}
