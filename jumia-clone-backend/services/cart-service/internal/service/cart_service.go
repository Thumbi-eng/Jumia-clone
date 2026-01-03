package service

import (
	"jumia-clone-backend/services/cart-service/internal/models"
	"jumia-clone-backend/services/cart-service/internal/repository"
)

type CartService interface {
	AddToCart(userID, productID, productName, imageURL string, quantity int, price float64) (*models.Cart, error)
	UpdateCartItem(userID, productID string, quantity int) (*models.Cart, error)
	RemoveFromCart(userID, productID string) (*models.Cart, error)
	GetCart(userID string) (*models.Cart, error)
	ClearCart(userID string) error
}

type cartService struct {
	repo repository.CartRepository
}

func NewCartService(repo repository.CartRepository) CartService {
	return &cartService{repo: repo}
}

func (s *cartService) AddToCart(userID, productID, productName, imageURL string, quantity int, price float64) (*models.Cart, error) {
	cart, err := s.repo.GetOrCreateCart(userID)
	if err != nil {
		return nil, err
	}

	_, err = s.repo.AddItem(cart.ID, productID, productName, imageURL, quantity, price)
	if err != nil {
		return nil, err
	}

	return s.repo.GetCart(userID)
}

func (s *cartService) UpdateCartItem(userID, productID string, quantity int) (*models.Cart, error) {
	cart, err := s.repo.GetOrCreateCart(userID)
	if err != nil {
		return nil, err
	}

	err = s.repo.UpdateItem(cart.ID, productID, quantity)
	if err != nil {
		return nil, err
	}

	return s.repo.GetCart(userID)
}

func (s *cartService) RemoveFromCart(userID, productID string) (*models.Cart, error) {
	cart, err := s.repo.GetOrCreateCart(userID)
	if err != nil {
		return nil, err
	}

	err = s.repo.RemoveItem(cart.ID, productID)
	if err != nil {
		return nil, err
	}

	return s.repo.GetCart(userID)
}

func (s *cartService) GetCart(userID string) (*models.Cart, error) {
	return s.repo.GetCart(userID)
}

func (s *cartService) ClearCart(userID string) error {
	return s.repo.ClearCart(userID)
}
