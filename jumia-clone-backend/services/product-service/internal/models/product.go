package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID                 string         `gorm:"type:uuid;primary_key" json:"id"`
	Name               string         `gorm:"type:varchar(255);not null" json:"name"`
	Description        string         `gorm:"type:text" json:"description"`
	Price              float64        `gorm:"type:decimal(10,2);not null" json:"price"`
	Category           string         `gorm:"type:varchar(100);index" json:"category"`
	Stock              int            `gorm:"default:0" json:"stock"`
	ImageURL           string         `gorm:"type:varchar(500)" json:"image_url"`
	Brand              string         `gorm:"type:varchar(100)" json:"brand"`
	DiscountPercentage float64        `gorm:"type:decimal(5,2);default:0" json:"discount_percentage"`
	IsActive           bool           `gorm:"default:true" json:"is_active"`
	IsFlashSale        bool           `gorm:"default:false;index" json:"is_flash_sale"`
	FlashSalePrice     float64        `gorm:"type:decimal(10,2);default:0" json:"flash_sale_price"`
	FlashSaleEndTime   *time.Time     `gorm:"type:timestamp" json:"flash_sale_end_time"`
	InitialStock       int            `gorm:"default:0" json:"initial_stock"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	return nil
}

func (Product) TableName() string {
	return "products"
}

func (p *Product) GetFinalPrice() float64 {
	if p.DiscountPercentage > 0 {
		return p.Price * (1 - p.DiscountPercentage/100)
	}
	return p.Price
}

func (p *Product) IsInStock() bool {
	return p.Stock > 0 && p.IsActive
}

func (p *Product) GetFlashSaleProgress() int {
	if p.InitialStock == 0 {
		return 0
	}
	sold := p.InitialStock - p.Stock
	return int((float64(sold) / float64(p.InitialStock)) * 100)
}

func (p *Product) IsFlashSaleActive() bool {
	if !p.IsFlashSale || p.FlashSaleEndTime == nil {
		return false
	}
	return p.FlashSaleEndTime.After(time.Now())
}
