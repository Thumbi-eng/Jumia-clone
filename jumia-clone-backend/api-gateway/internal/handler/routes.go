package handler

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// RegisterAllRoutes registers all routes for all microservices
func RegisterAllRoutes(router *gin.Engine, userConn, productConn, cartConn, orderConn *grpc.ClientConn) {
	// Initialize handlers
	userHandler := NewUserHandler(userConn)
	productHandler := NewProductHandler(productConn)
	cartHandler := NewCartHandler(cartConn)
	orderHandler := NewOrderHandler(orderConn)

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// User routes
		users := v1.Group("/users")
		{
			users.POST("/register", userHandler.Register)
			users.POST("/login", userHandler.Login)
			users.GET("/:id", userHandler.AuthMiddleware(), userHandler.GetUser)
			users.PUT("/:id", userHandler.AuthMiddleware(), userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.AuthMiddleware(), userHandler.DeleteUser)
		}

		// Auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/verify", userHandler.VerifyToken)
		}

		// Product routes
		products := v1.Group("/products")
		{
			products.POST("", productHandler.CreateProduct)
			products.GET("", productHandler.ListProducts)
			products.GET("/search", productHandler.SearchProducts)
			products.GET("/category", productHandler.GetProductsByCategory)
			products.GET("/:id", productHandler.GetProduct)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
		}

		// Cart routes
		cart := v1.Group("/cart")
		{
			cart.POST("/add", cartHandler.AddToCart)
			cart.PUT("/update", cartHandler.UpdateCartItem)
			cart.POST("/remove", cartHandler.RemoveFromCart)
			cart.GET("/:user_id", cartHandler.GetCart)
			cart.DELETE("/:user_id", cartHandler.ClearCart)
		}

		// Order routes
		orders := v1.Group("/orders")
		{
			orders.POST("", orderHandler.CreateOrder)
			orders.GET("", orderHandler.ListOrders)
			orders.GET("/:id", orderHandler.GetOrder)
			orders.PUT("/:id/status", orderHandler.UpdateOrderStatus)
			orders.POST("/:id/cancel", orderHandler.CancelOrder)
		}
	}
}
