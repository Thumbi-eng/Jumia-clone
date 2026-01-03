package repository

import (
	"errors"

	"jumia-clone-backend/services/order-service/internal/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *models.Order) error
	GetOrder(orderID string) (*models.Order, error)
	ListOrders(userID string, page, pageSize int) ([]models.Order, int64, error)
	UpdateOrderStatus(orderID, status string) error
	CancelOrder(orderID, userID string) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) GetOrder(orderID string) (*models.Order, error) {
	var order models.Order
	err := r.db.Where("id = ?", orderID).Preload("Items").First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) ListOrders(userID string, page, pageSize int) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64

	query := r.db.Where("user_id = ?", userID)

	if err := query.Model(&models.Order{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := query.Preload("Items").
		Order("created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&orders).Error
	if err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}

func (r *orderRepository) UpdateOrderStatus(orderID, status string) error {
	result := r.db.Model(&models.Order{}).Where("id = ?", orderID).Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("order not found")
	}
	return nil
}

func (r *orderRepository) CancelOrder(orderID, userID string) error {
	var order models.Order
	err := r.db.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("order not found or unauthorized")
		}
		return err
	}

	if order.Status == "delivered" || order.Status == "cancelled" {
		return errors.New("cannot cancel order with status: " + order.Status)
	}

	return r.db.Model(&order).Update("status", "cancelled").Error
}
