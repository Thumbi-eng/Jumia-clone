package handler

import (
	"context"
	"time"

	"jumia-clone-backend/services/product-service/internal/models"
	"jumia-clone-backend/services/product-service/internal/service"
	pb "jumia-clone-backend/services/product-service/proto"
)

// ProductServiceHandler implements the gRPC ProductServiceServer
type ProductServiceHandler struct {
	pb.UnimplementedProductServiceServer
	productService service.ProductService
}

// NewProductServiceHandler creates a new product service handler
func NewProductServiceHandler(productService service.ProductService) *ProductServiceHandler {
	return &ProductServiceHandler{
		productService: productService,
	}
}

// CreateProduct handles product creation
func (h *ProductServiceHandler) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product, err := h.productService.CreateProduct(
		req.Name,
		req.Description,
		req.Category,
		req.ImageUrl,
		req.Brand,
		req.Price,
		req.DiscountPercentage,
		int(req.Stock),
	)
	if err != nil {
		return &pb.CreateProductResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.CreateProductResponse{
		Id:      product.ID,
		Success: true,
		Message: "Product created successfully",
	}, nil
}

// GetProduct retrieves product information
func (h *ProductServiceHandler) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	product, err := h.productService.GetProductByID(req.Id)
	if err != nil {
		return &pb.GetProductResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.GetProductResponse{
		Success: true,
		Message: "Product retrieved successfully",
		Product: convertToProductData(product),
	}, nil
}

// UpdateProduct handles product updates
func (h *ProductServiceHandler) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	product, err := h.productService.UpdateProduct(
		req.Id,
		req.Name,
		req.Description,
		req.Category,
		req.ImageUrl,
		req.Brand,
		req.Price,
		req.DiscountPercentage,
		int(req.Stock),
	)
	if err != nil {
		return &pb.UpdateProductResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.UpdateProductResponse{
		Success: true,
		Message: "Product updated successfully",
		Product: convertToProductData(product),
	}, nil
}

// DeleteProduct handles product deletion
func (h *ProductServiceHandler) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	err := h.productService.DeleteProduct(req.Id)
	if err != nil {
		return &pb.DeleteProductResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.DeleteProductResponse{
		Success: true,
		Message: "Product deleted successfully",
	}, nil
}

// ListProducts retrieves paginated products
func (h *ProductServiceHandler) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	products, total, err := h.productService.ListProducts(int(req.Page), int(req.PageSize))
	if err != nil {
		return &pb.ListProductsResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.ListProductsResponse{
		Success:  true,
		Message:  "Products retrieved successfully",
		Products: convertToProductDataList(products),
		Total:    int32(total),
	}, nil
}

// SearchProducts searches for products
func (h *ProductServiceHandler) SearchProducts(ctx context.Context, req *pb.SearchProductsRequest) (*pb.SearchProductsResponse, error) {
	products, total, err := h.productService.SearchProducts(req.Query, int(req.Page), int(req.PageSize))
	if err != nil {
		return &pb.SearchProductsResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.SearchProductsResponse{
		Success:  true,
		Message:  "Search completed successfully",
		Products: convertToProductDataList(products),
		Total:    int32(total),
	}, nil
}

// GetProductsByCategory retrieves products by category
func (h *ProductServiceHandler) GetProductsByCategory(ctx context.Context, req *pb.GetProductsByCategoryRequest) (*pb.GetProductsByCategoryResponse, error) {
	products, total, err := h.productService.GetProductsByCategory(req.Category, int(req.Page), int(req.PageSize))
	if err != nil {
		return &pb.GetProductsByCategoryResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.GetProductsByCategoryResponse{
		Success:  true,
		Message:  "Products retrieved successfully",
		Products: convertToProductDataList(products),
		Total:    int32(total),
	}, nil
}

// Helper function to convert model to protobuf
func convertToProductData(product *models.Product) *pb.ProductData {
	return &pb.ProductData{
		Id:                 product.ID,
		Name:               product.Name,
		Description:        product.Description,
		Price:              product.Price,
		Category:           product.Category,
		Stock:              int32(product.Stock),
		ImageUrl:           product.ImageURL,
		Brand:              product.Brand,
		DiscountPercentage: product.DiscountPercentage,
		FinalPrice:         product.GetFinalPrice(),
		InStock:            product.IsInStock(),
		CreatedAt:          product.CreatedAt.Format(time.RFC3339),
		UpdatedAt:          product.UpdatedAt.Format(time.RFC3339),
	}
}

// Helper function to convert model list to protobuf list
func convertToProductDataList(products []*models.Product) []*pb.ProductData {
	var productDataList []*pb.ProductData
	for _, product := range products {
		productDataList = append(productDataList, convertToProductData(product))
	}
	return productDataList
}
