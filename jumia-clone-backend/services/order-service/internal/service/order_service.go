package service

import (
	"jumia-clone-backend/services/order-service/internal/models"
	"jumia-clone-backend/services/order-service/internal/repository"
)

type OrderService interface {
	CreateOrder(userID, shippingAddress, paymentMethod string, items []OrderItemInput) (*models.Order, error)
	GetOrder(orderID string) (*models.Order, error)
	ListOrders(userID string, page, pageSize int) ([]models.Order, int64, error)
	UpdateOrderStatus(orderID, status string) (*models.Order, error)
	CancelOrder(orderID, userID string) error
}

type OrderItemInput struct {
	ProductID   string
	ProductName string
	Quantity    int
	Price       float64
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{repo: repo}
}

func (s *orderService) CreateOrder(userID, shippingAddress, paymentMethod string, items []OrderItemInput) (*models.Order, error) {
	orderItems := make([]models.OrderItem, 0, len(items))
	var totalPrice float64

	for _, item := range items {
		orderItem := models.OrderItem{
			ProductID:   item.ProductID,
			ProductName: item.ProductName,
			Quantity:    item.Quantity,
			Price:       item.Price,
		}
		orderItems = append(orderItems, orderItem)
		totalPrice += item.Price * float64(item.Quantity)
	}

	order := &models.Order{
		UserID:          userID,
		Items:           orderItems,
		Status:          "pending",
		TotalPrice:      totalPrice,
		ShippingAddress: shippingAddress,
		PaymentMethod:   paymentMethod,
	}

	if err := s.repo.CreateOrder(order); err != nil {
		return nil, err
	}

	return s.repo.GetOrder(order.ID)
}

func (s *orderService) GetOrder(orderID string) (*models.Order, error) {
	return s.repo.GetOrder(orderID)
}

func (s *orderService) ListOrders(userID string, page, pageSize int) ([]models.Order, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	return s.repo.ListOrders(userID, page, pageSize)
}

func (s *orderService) UpdateOrderStatus(orderID, status string) (*models.Order, error) {
	if err := s.repo.UpdateOrderStatus(orderID, status); err != nil {
		return nil, err
	}
	return s.repo.GetOrder(orderID)
}

func (s *orderService) CancelOrder(orderID, userID string) error {
	return s.repo.CancelOrder(orderID, userID)
}
