package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	ID        string     `gorm:"type:uuid;primary_key" json:"id"`
	UserID    string     `gorm:"type:uuid;uniqueIndex;not null" json:"user_id"`
	Items     []CartItem `gorm:"foreignKey:CartID;constraint:OnDelete:CASCADE" json:"items"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type CartItem struct {
	ID          string    `gorm:"type:uuid;primary_key" json:"id"`
	CartID      string    `gorm:"type:uuid;not null;index" json:"cart_id"`
	ProductID   string    `gorm:"type:uuid;not null" json:"product_id"`
	ProductName string    `gorm:"type:varchar(255)" json:"product_name"`
	Quantity    int       `gorm:"not null;default:1" json:"quantity"`
	Price       float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	ImageURL    string    `gorm:"type:varchar(500)" json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Cart) TableName() string {
	return "carts"
}

func (CartItem) TableName() string {
	return "cart_items"
}

func (c *Cart) BeforeCreate(tx *gorm.DB) error {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return nil
}

func (ci *CartItem) BeforeCreate(tx *gorm.DB) error {
	if ci.ID == "" {
		ci.ID = uuid.New().String()
	}
	return nil
}

func (ci *CartItem) GetSubtotal() float64 {
	return ci.Price * float64(ci.Quantity)
}

func (c *Cart) GetTotalPrice() float64 {
	var total float64
	for _, item := range c.Items {
		total += item.GetSubtotal()
	}
	return total
}

func (c *Cart) GetTotalItems() int {
	var total int
	for _, item := range c.Items {
		total += item.Quantity
	}
	return total
}
