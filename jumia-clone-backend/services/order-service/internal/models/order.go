package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID              string      `gorm:"type:uuid;primary_key" json:"id"`
	UserID          string      `gorm:"type:uuid;not null;index" json:"user_id"`
	Items           []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE" json:"items"`
	Status          string      `gorm:"type:varchar(50);not null;default:'pending'" json:"status"`
	TotalPrice      float64     `gorm:"type:decimal(10,2);not null" json:"total_price"`
	ShippingAddress string      `gorm:"type:text" json:"shipping_address"`
	PaymentMethod   string      `gorm:"type:varchar(50)" json:"payment_method"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

type OrderItem struct {
	ID          string    `gorm:"type:uuid;primary_key" json:"id"`
	OrderID     string    `gorm:"type:uuid;not null;index" json:"order_id"`
	ProductID   string    `gorm:"type:uuid;not null" json:"product_id"`
	ProductName string    `gorm:"type:varchar(255)" json:"product_name"`
	Quantity    int       `gorm:"not null" json:"quantity"`
	Price       float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Order) TableName() string {
	return "orders"
}

func (OrderItem) TableName() string {
	return "order_items"
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	if o.ID == "" {
		o.ID = uuid.New().String()
	}
	return nil
}

func (oi *OrderItem) BeforeCreate(tx *gorm.DB) error {
	if oi.ID == "" {
		oi.ID = uuid.New().String()
	}
	return nil
}

func (oi *OrderItem) GetSubtotal() float64 {
	return oi.Price * float64(oi.Quantity)
}

func (o *Order) CalculateTotalPrice() float64 {
	var total float64
	for _, item := range o.Items {
		total += item.GetSubtotal()
	}
	return total
}
