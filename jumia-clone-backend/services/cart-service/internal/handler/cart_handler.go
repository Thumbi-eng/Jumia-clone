package handler

import (
	"context"
	"time"

	"jumia-clone-backend/services/cart-service/internal/models"
	"jumia-clone-backend/services/cart-service/internal/service"
	pb "jumia-clone-backend/services/cart-service/proto"
)

type CartServiceHandler struct {
	pb.UnimplementedCartServiceServer
	cartService service.CartService
}

func NewCartHandler(cartService service.CartService) *CartServiceHandler {
	return &CartServiceHandler{
		cartService: cartService,
	}
}

func (h *CartServiceHandler) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.AddToCartResponse, error) {
	cart, err := h.cartService.AddToCart(
		req.UserId,
		req.ProductId,
		req.ProductName,
		req.ImageUrl,
		int(req.Quantity),
		req.Price,
	)
	if err != nil {
		return &pb.AddToCartResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.AddToCartResponse{
		Success: true,
		Message: "Item added to cart successfully",
		Cart:    convertToCartData(cart),
	}, nil
}

func (h *CartServiceHandler) UpdateCartItem(ctx context.Context, req *pb.UpdateCartItemRequest) (*pb.UpdateCartItemResponse, error) {
	cart, err := h.cartService.UpdateCartItem(req.UserId, req.ProductId, int(req.Quantity))
	if err != nil {
		return &pb.UpdateCartItemResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.UpdateCartItemResponse{
		Success: true,
		Message: "Cart item updated successfully",
		Cart:    convertToCartData(cart),
	}, nil
}

func (h *CartServiceHandler) RemoveFromCart(ctx context.Context, req *pb.RemoveFromCartRequest) (*pb.RemoveFromCartResponse, error) {
	cart, err := h.cartService.RemoveFromCart(req.UserId, req.ProductId)
	if err != nil {
		return &pb.RemoveFromCartResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.RemoveFromCartResponse{
		Success: true,
		Message: "Item removed from cart successfully",
		Cart:    convertToCartData(cart),
	}, nil
}

func (h *CartServiceHandler) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	cart, err := h.cartService.GetCart(req.UserId)
	if err != nil {
		return &pb.GetCartResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.GetCartResponse{
		Success: true,
		Message: "Cart retrieved successfully",
		Cart:    convertToCartData(cart),
	}, nil
}

func (h *CartServiceHandler) ClearCart(ctx context.Context, req *pb.ClearCartRequest) (*pb.ClearCartResponse, error) {
	err := h.cartService.ClearCart(req.UserId)
	if err != nil {
		return &pb.ClearCartResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.ClearCartResponse{
		Success: true,
		Message: "Cart cleared successfully",
	}, nil
}

func convertToCartData(cart *models.Cart) *pb.CartData {
	items := make([]*pb.CartItemData, 0, len(cart.Items))
	for _, item := range cart.Items {
		items = append(items, &pb.CartItemData{
			Id:          item.ID,
			ProductId:   item.ProductID,
			ProductName: item.ProductName,
			ImageUrl:    item.ImageURL,
			Quantity:    int32(item.Quantity),
			Price:       item.Price,
			Subtotal:    item.GetSubtotal(),
		})
	}

	return &pb.CartData{
		Id:         cart.ID,
		UserId:     cart.UserID,
		Items:      items,
		TotalPrice: cart.GetTotalPrice(),
		TotalItems: int32(cart.GetTotalItems()),
		CreatedAt:  cart.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  cart.UpdatedAt.Format(time.RFC3339),
	}
}
