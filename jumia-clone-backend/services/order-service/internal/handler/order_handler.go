package handler

import (
	"context"
	"time"

	"jumia-clone-backend/services/order-service/internal/models"
	"jumia-clone-backend/services/order-service/internal/service"
	pb "jumia-clone-backend/services/order-service/proto"
)

type OrderServiceHandler struct {
	pb.UnimplementedOrderServiceServer
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) *OrderServiceHandler {
	return &OrderServiceHandler{
		orderService: orderService,
	}
}

func (h *OrderServiceHandler) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	items := make([]service.OrderItemInput, 0, len(req.Items))
	for _, item := range req.Items {
		items = append(items, service.OrderItemInput{
			ProductID:   item.ProductId,
			ProductName: item.ProductName,
			Quantity:    int(item.Quantity),
			Price:       item.Price,
		})
	}

	order, err := h.orderService.CreateOrder(req.UserId, req.ShippingAddress, req.PaymentMethod, items)
	if err != nil {
		return &pb.CreateOrderResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.CreateOrderResponse{
		Success: true,
		Message: "Order created successfully",
		Order:   convertToOrderData(order),
	}, nil
}

func (h *OrderServiceHandler) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	order, err := h.orderService.GetOrder(req.OrderId)
	if err != nil {
		return &pb.GetOrderResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.GetOrderResponse{
		Success: true,
		Message: "Order retrieved successfully",
		Order:   convertToOrderData(order),
	}, nil
}

func (h *OrderServiceHandler) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	page := int(req.Page)
	pageSize := int(req.PageSize)
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	orders, total, err := h.orderService.ListOrders(req.UserId, page, pageSize)
	if err != nil {
		return &pb.ListOrdersResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	orderDataList := make([]*pb.OrderData, 0, len(orders))
	for _, order := range orders {
		orderDataList = append(orderDataList, convertToOrderData(&order))
	}

	return &pb.ListOrdersResponse{
		Success: true,
		Message: "Orders retrieved successfully",
		Orders:  orderDataList,
		Total:   int32(total),
	}, nil
}

func (h *OrderServiceHandler) UpdateOrderStatus(ctx context.Context, req *pb.UpdateOrderStatusRequest) (*pb.UpdateOrderStatusResponse, error) {
	order, err := h.orderService.UpdateOrderStatus(req.OrderId, req.Status)
	if err != nil {
		return &pb.UpdateOrderStatusResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.UpdateOrderStatusResponse{
		Success: true,
		Message: "Order status updated successfully",
		Order:   convertToOrderData(order),
	}, nil
}

func (h *OrderServiceHandler) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	err := h.orderService.CancelOrder(req.OrderId, req.UserId)
	if err != nil {
		return &pb.CancelOrderResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.CancelOrderResponse{
		Success: true,
		Message: "Order cancelled successfully",
	}, nil
}

func convertToOrderData(order *models.Order) *pb.OrderData {
	items := make([]*pb.OrderItemData, 0, len(order.Items))
	for _, item := range order.Items {
		items = append(items, &pb.OrderItemData{
			Id:          item.ID,
			ProductId:   item.ProductID,
			ProductName: item.ProductName,
			Quantity:    int32(item.Quantity),
			Price:       item.Price,
			Subtotal:    item.GetSubtotal(),
		})
	}

	return &pb.OrderData{
		Id:              order.ID,
		UserId:          order.UserID,
		Items:           items,
		Status:          order.Status,
		TotalPrice:      order.TotalPrice,
		ShippingAddress: order.ShippingAddress,
		PaymentMethod:   order.PaymentMethod,
		CreatedAt:       order.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       order.UpdatedAt.Format(time.RFC3339),
	}
}
