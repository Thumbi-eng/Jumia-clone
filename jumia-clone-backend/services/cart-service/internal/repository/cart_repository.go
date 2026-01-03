package repository

import (
	"errors"

	"jumia-clone-backend/services/cart-service/internal/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	GetOrCreateCart(userID string) (*models.Cart, error)
	AddItem(cartID, productID, productName, imageURL string, quantity int, price float64) (*models.CartItem, error)
	UpdateItem(cartID, productID string, quantity int) error
	RemoveItem(cartID, productID string) error
	GetCart(userID string) (*models.Cart, error)
	ClearCart(userID string) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) GetOrCreateCart(userID string) (*models.Cart, error) {
	var cart models.Cart
	err := r.db.Where("user_id = ?", userID).Preload("Items").First(&cart).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cart = models.Cart{UserID: userID}
			if err := r.db.Create(&cart).Error; err != nil {
				return nil, err
			}
			return &cart, nil
		}
		return nil, err
	}
	return &cart, nil
}

func (r *cartRepository) AddItem(cartID, productID, productName, imageURL string, quantity int, price float64) (*models.CartItem, error) {
	var existingItem models.CartItem
	err := r.db.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&existingItem).Error

	if err == nil {
		existingItem.Quantity += quantity
		if err := r.db.Save(&existingItem).Error; err != nil {
			return nil, err
		}
		return &existingItem, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	item := models.CartItem{
		CartID:      cartID,
		ProductID:   productID,
		ProductName: productName,
		ImageURL:    imageURL,
		Quantity:    quantity,
		Price:       price,
	}

	if err := r.db.Create(&item).Error; err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *cartRepository) UpdateItem(cartID, productID string, quantity int) error {
	var item models.CartItem
	err := r.db.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&item).Error
	if err != nil {
		return err
	}

	item.Quantity = quantity
	return r.db.Save(&item).Error
}

func (r *cartRepository) RemoveItem(cartID, productID string) error {
	return r.db.Where("cart_id = ? AND product_id = ?", cartID, productID).Delete(&models.CartItem{}).Error
}

func (r *cartRepository) GetCart(userID string) (*models.Cart, error) {
	var cart models.Cart
	err := r.db.Where("user_id = ?", userID).Preload("Items").First(&cart).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &models.Cart{UserID: userID, Items: []models.CartItem{}}, nil
		}
		return nil, err
	}
	return &cart, nil
}

func (r *cartRepository) ClearCart(userID string) error {
	var cart models.Cart
	err := r.db.Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	return r.db.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error
}
