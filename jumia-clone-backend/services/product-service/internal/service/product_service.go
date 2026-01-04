package service

import (
	"fmt"
	"time"

	"jumia-clone-backend/services/product-service/internal/models"
	"jumia-clone-backend/services/product-service/internal/repository"
)

type ProductService interface {
	CreateProduct(name, description, category, imageURL, brand string, price, discountPercentage float64, stock int, isFlashSale bool, flashSalePrice float64, flashSaleEndTime *string, initialStock int, isTopDeal bool, dealType string, dealPriority int) (*models.Product, error)
	GetProductByID(id string) (*models.Product, error)
	UpdateProduct(id, name, description, category, imageURL, brand string, price, discountPercentage float64, stock int, isFlashSale bool, flashSalePrice float64, flashSaleEndTime *string, initialStock int, isTopDeal bool, dealType string, dealPriority int) (*models.Product, error)
	DeleteProduct(id string) error
	ListProducts(page, pageSize int) ([]*models.Product, int64, error)
	SearchProducts(query string, page, pageSize int) ([]*models.Product, int64, error)
	GetProductsByCategory(category string, page, pageSize int) ([]*models.Product, int64, error)
	GetTopDeals(page, pageSize int) ([]*models.Product, int64, error)
	GetDealsByType(dealType string, page, pageSize int) ([]*models.Product, int64, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(name, description, category, imageURL, brand string, price, discountPercentage float64, stock int, isFlashSale bool, flashSalePrice float64, flashSaleEndTime *string, initialStock int, isTopDeal bool, dealType string, dealPriority int) (*models.Product, error) {
	product := &models.Product{
		Name:               name,
		Description:        description,
		Price:              price,
		Category:           category,
		Stock:              stock,
		ImageURL:           imageURL,
		Brand:              brand,
		DiscountPercentage: discountPercentage,
		IsActive:           true,
		IsFlashSale:        isFlashSale,
		FlashSalePrice:     flashSalePrice,
		InitialStock:       initialStock,
		IsTopDeal:          isTopDeal,
		DealType:           dealType,
		DealPriority:       dealPriority,
	}

	// Parse flash sale end time if provided
	if flashSaleEndTime != nil && *flashSaleEndTime != "" {
		if endTime, err := parseTime(*flashSaleEndTime); err == nil {
			product.FlashSaleEndTime = &endTime
		}
	}

	if err := s.repo.Create(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) GetProductByID(id string) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *productService) UpdateProduct(id, name, description, category, imageURL, brand string, price, discountPercentage float64, stock int, isFlashSale bool, flashSalePrice float64, flashSaleEndTime *string, initialStock int, isTopDeal bool, dealType string, dealPriority int) (*models.Product, error) {
	product, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if name != "" {
		product.Name = name
	}
	if description != "" {
		product.Description = description
	}
	if category != "" {
		product.Category = category
	}
	if imageURL != "" {
		product.ImageURL = imageURL
	}
	if brand != "" {
		product.Brand = brand
	}
	if price > 0 {
		product.Price = price
	}
	if discountPercentage >= 0 {
		product.DiscountPercentage = discountPercentage
	}
	if stock >= 0 {
		product.Stock = stock
	}

	// Update flash sale fields
	product.IsFlashSale = isFlashSale
	product.FlashSalePrice = flashSalePrice
	if initialStock > 0 {
		product.InitialStock = initialStock
	}

	// Update deal fields
	product.IsTopDeal = isTopDeal
	product.DealType = dealType
	product.DealPriority = dealPriority

	// Parse flash sale end time if provided
	if flashSaleEndTime != nil && *flashSaleEndTime != "" {
		if endTime, err := parseTime(*flashSaleEndTime); err == nil {
			product.FlashSaleEndTime = &endTime
		}
	} else if !isFlashSale {
		// Clear flash sale end time if not a flash sale
		product.FlashSaleEndTime = nil
	}

	if err := s.repo.Update(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) DeleteProduct(id string) error {
	return s.repo.Delete(id)
}

func (s *productService) ListProducts(page, pageSize int) ([]*models.Product, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	return s.repo.List(page, pageSize)
}

func (s *productService) SearchProducts(query string, page, pageSize int) ([]*models.Product, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	return s.repo.Search(query, page, pageSize)
}

func (s *productService) GetProductsByCategory(category string, page, pageSize int) ([]*models.Product, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	return s.repo.GetByCategory(category, page, pageSize)
}

func (s *productService) GetTopDeals(page, pageSize int) ([]*models.Product, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	return s.repo.GetTopDeals(page, pageSize)
}

func (s *productService) GetDealsByType(dealType string, page, pageSize int) ([]*models.Product, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	return s.repo.GetDealsByType(dealType, page, pageSize)
}

// Helper function to parse time string
func parseTime(timeStr string) (time.Time, error) {
	// Try parsing ISO 8601 format first
	if t, err := time.Parse(time.RFC3339, timeStr); err == nil {
		return t, nil
	}
	// Try other common formats
	formats := []string{
		time.RFC3339Nano,
		"2006-01-02T15:04:05Z",
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05",
	}
	for _, format := range formats {
		if t, err := time.Parse(format, timeStr); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unable to parse time: %s", timeStr)
}
