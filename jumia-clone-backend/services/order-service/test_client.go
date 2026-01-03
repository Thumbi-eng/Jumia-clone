package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "jumia-clone-backend/services/order-service/proto"
)

const (
	address        = "localhost:50054"
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

	client := pb.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	log.Println("===== Order Service Client Test =====\n")

	// Test 1: Create order
	log.Println("Test 1: Creating order with 2 items...")
	createResp, err := client.CreateOrder(ctx, &pb.CreateOrderRequest{
		UserId: testUserID,
		Items: []*pb.OrderItemInput{
			{
				ProductId:   testProductID1,
				ProductName: "Smartphone",
				Quantity:    2,
				Price:       299.99,
			},
			{
				ProductId:   testProductID2,
				ProductName: "Headphones",
				Quantity:    1,
				Price:       49.99,
			},
		},
		ShippingAddress: "123 Main St, Nairobi, Kenya",
		PaymentMethod:   "Credit Card",
	})
	if err != nil {
		log.Fatalf("CreateOrder failed: %v", err)
	}
	log.Printf("Success: %v, Message: %s\n", createResp.Success, createResp.Message)
	log.Printf("Order ID: %s, Status: %s, Total: $%.2f\n", createResp.Order.Id, createResp.Order.Status, createResp.Order.TotalPrice)
	log.Printf("Shipping: %s, Payment: %s\n", createResp.Order.ShippingAddress, createResp.Order.PaymentMethod)
	log.Println("Items:")
	for i, item := range createResp.Order.Items {
		log.Printf("  %d. %s (x%d) - $%.2f each, Subtotal: $%.2f\n", i+1, item.ProductName, item.Quantity, item.Price, item.Subtotal)
	}
	log.Println()

	orderID := createResp.Order.Id

	// Test 2: Get order
	log.Println("Test 2: Getting order details...")
	getResp, err := client.GetOrder(ctx, &pb.GetOrderRequest{
		OrderId: orderID,
	})
	if err != nil {
		log.Fatalf("GetOrder failed: %v", err)
	}
	log.Printf("Success: %v, Message: %s\n", getResp.Success, getResp.Message)
	log.Printf("Order ID: %s, Status: %s, Total: $%.2f\n", getResp.Order.Id, getResp.Order.Status, getResp.Order.TotalPrice)
	log.Println()

	// Test 3: Create another order
	log.Println("Test 3: Creating second order...")
	createResp2, err := client.CreateOrder(ctx, &pb.CreateOrderRequest{
		UserId: testUserID,
		Items: []*pb.OrderItemInput{
			{
				ProductId:   testProductID1,
				ProductName: "Smartphone",
				Quantity:    1,
				Price:       299.99,
			},
		},
		ShippingAddress: "456 Oak Ave, Mombasa, Kenya",
		PaymentMethod:   "M-Pesa",
	})
	if err != nil {
		log.Fatalf("CreateOrder failed: %v", err)
	}
	log.Printf("Success: %v, Message: %s\n", createResp2.Success, createResp2.Message)
	log.Printf("Order ID: %s, Status: %s, Total: $%.2f\n\n", createResp2.Order.Id, createResp2.Order.Status, createResp2.Order.TotalPrice)

	orderID2 := createResp2.Order.Id

	// Test 4: List user orders
	log.Println("Test 4: Listing all user orders...")
	listResp, err := client.ListOrders(ctx, &pb.ListOrdersRequest{
		UserId:   testUserID,
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		log.Fatalf("ListOrders failed: %v", err)
	}
	log.Printf("Success: %v, Message: %s\n", listResp.Success, listResp.Message)
	log.Printf("Total orders: %d\n", listResp.Total)
	for i, order := range listResp.Orders {
		log.Printf("  %d. Order %s - Status: %s, Items: %d, Total: $%.2f\n",
			i+1, order.Id[:8]+"...", order.Status, len(order.Items), order.TotalPrice)
	}
	log.Println()

	// Test 5: Update order status
	log.Println("Test 5: Updating first order status to 'processing'...")
	updateResp, err := client.UpdateOrderStatus(ctx, &pb.UpdateOrderStatusRequest{
		OrderId: orderID,
		Status:  "processing",
	})
	if err != nil {
		log.Fatalf("UpdateOrderStatus failed: %v", err)
	}
	log.Printf("Success: %v, Message: %s\n", updateResp.Success, updateResp.Message)
	log.Printf("Order ID: %s, New Status: %s\n\n", updateResp.Order.Id[:8]+"...", updateResp.Order.Status)

	// Test 6: Update to shipped
	log.Println("Test 6: Updating first order status to 'shipped'...")
	updateResp2, err := client.UpdateOrderStatus(ctx, &pb.UpdateOrderStatusRequest{
		OrderId: orderID,
		Status:  "shipped",
	})
	if err != nil {
		log.Fatalf("UpdateOrderStatus failed: %v", err)
	}
	log.Printf("Success: %v, Message: %s\n", updateResp2.Success, updateResp2.Message)
	log.Printf("Order ID: %s, New Status: %s\n\n", updateResp2.Order.Id[:8]+"...", updateResp2.Order.Status)

	// Test 7: Cancel second order
	log.Println("Test 7: Cancelling second order...")
	cancelResp, err := client.CancelOrder(ctx, &pb.CancelOrderRequest{
		OrderId: orderID2,
		UserId:  testUserID,
	})
	if err != nil {
		log.Fatalf("CancelOrder failed: %v", err)
	}
	log.Printf("Success: %v, Message: %s\n\n", cancelResp.Success, cancelResp.Message)

	// Test 8: Verify cancellation
	log.Println("Test 8: Verifying cancelled order status...")
	getResp2, err := client.GetOrder(ctx, &pb.GetOrderRequest{
		OrderId: orderID2,
	})
	if err != nil {
		log.Fatalf("GetOrder failed: %v", err)
	}
	log.Printf("Order ID: %s, Status: %s\n\n", getResp2.Order.Id[:8]+"...", getResp2.Order.Status)

	// Test 9: List final orders
	log.Println("Test 9: Listing final user orders...")
	listResp2, err := client.ListOrders(ctx, &pb.ListOrdersRequest{
		UserId:   testUserID,
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		log.Fatalf("ListOrders failed: %v", err)
	}
	log.Printf("Total orders: %d\n", listResp2.Total)
	for i, order := range listResp2.Orders {
		log.Printf("  %d. Order %s - Status: %s, Total: $%.2f, Created: %s\n",
			i+1, order.Id[:8]+"...", order.Status, order.TotalPrice, order.CreatedAt)
	}

	log.Println("\n===== All Tests Completed Successfully! =====")
}
