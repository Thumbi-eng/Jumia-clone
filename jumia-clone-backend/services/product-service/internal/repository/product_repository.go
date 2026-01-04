package repository

import (
	"errors"

	"jumia-clone-backend/services/product-service/internal/models"

	"gorm.io/gorm"
)

// ProductRepository defines methods for product data access
type ProductRepository interface {
	Create(product *models.Product) error
	GetByID(id string) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id string) error
	List(page, pageSize int) ([]*models.Product, int64, error)
	Search(query string, page, pageSize int) ([]*models.Product, int64, error)
	GetByCategory(category string, page, pageSize int) ([]*models.Product, int64, error)
	GetFlashSaleProducts(page, pageSize int) ([]*models.Product, int64, error)
}

type productRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new product repository
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

// Create creates a new product in the database
func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

// GetByID retrieves a product by ID
func (r *productRepository) GetByID(id string) (*models.Product, error) {
	var product models.Product
	err := r.db.Where("id = ? AND is_active = ?", id, true).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}

// Update updates an existing product
func (r *productRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

// Delete soft deletes a product (sets is_active to false)
func (r *productRepository) Delete(id string) error {
	return r.db.Model(&models.Product{}).Where("id = ?", id).Update("is_active", false).Error
}

// List retrieves paginated products
func (r *productRepository) List(page, pageSize int) ([]*models.Product, int64, error) {
	var products []*models.Product
	var total int64

	offset := (page - 1) * pageSize

	// Count total
	if err := r.db.Model(&models.Product{}).Where("is_active = ?", true).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	err := r.db.Where("is_active = ?", true).
		Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&products).Error

	return products, total, err
}

// Search searches products by name or description
func (r *productRepository) Search(query string, page, pageSize int) ([]*models.Product, int64, error) {
	var products []*models.Product
	var total int64

	offset := (page - 1) * pageSize
	searchPattern := "%" + query + "%"

	// Count total
	if err := r.db.Model(&models.Product{}).
		Where("is_active = ? AND (name ILIKE ? OR description ILIKE ?)", true, searchPattern, searchPattern).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	err := r.db.Where("is_active = ? AND (name ILIKE ? OR description ILIKE ?)", true, searchPattern, searchPattern).
		Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&products).Error

	return products, total, err
}

// GetByCategory retrieves products by category
func (r *productRepository) GetByCategory(category string, page, pageSize int) ([]*models.Product, int64, error) {
	var products []*models.Product
	var total int64

	offset := (page - 1) * pageSize

	// Count total
	if err := r.db.Model(&models.Product{}).
		Where("is_active = ? AND category = ?", true, category).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	err := r.db.Where("is_active = ? AND category = ?", true, category).
		Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&products).Error

	return products, total, err
}

// GetFlashSaleProducts retrieves active flash sale products
func (r *productRepository) GetFlashSaleProducts(page, pageSize int) ([]*models.Product, int64, error) {
	var products []*models.Product
	var total int64

	offset := (page - 1) * pageSize

	// Count total - only active flash sales
	if err := r.db.Model(&models.Product{}).
		Where("is_active = ? AND is_flash_sale = ? AND flash_sale_end_time > NOW()", true, true).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results - order by flash sale end time (ending soon first)
	err := r.db.Where("is_active = ? AND is_flash_sale = ? AND flash_sale_end_time > NOW()", true, true).
		Offset(offset).
		Limit(pageSize).
		Order("flash_sale_end_time ASC").
		Find(&products).Error

	return products, total, err
}
