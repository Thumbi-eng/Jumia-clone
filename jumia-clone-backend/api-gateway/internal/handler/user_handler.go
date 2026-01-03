package handler

import (
	"context"
	"net/http"
	"strings"
	"time"

	pb "jumia-clone-backend/api-gateway/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	userClient pb.UserServiceClient
}

// NewUserHandler creates a new user handler
func NewUserHandler(conn *grpc.ClientConn) *UserHandler {
	return &UserHandler{
		userClient: pb.NewUserServiceClient(conn),
	}
}

// RegisterRoutes registers all routes for the API gateway
func RegisterRoutes(router *gin.Engine, userConn *grpc.ClientConn) {
	userHandler := NewUserHandler(userConn)

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
	}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})
}

// Register handles user registration
func (h *UserHandler) Register(c *gin.Context) {
	var req struct {
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required,min=6"`
		Phone     string `json:"phone"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.userClient.Register(ctx, &pb.RegisterRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Phone:     req.Phone,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	if !resp.Success {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp.Message})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": resp.Message,
		"user_id": resp.UserId,
	})
}

// Login handles user authentication
func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.userClient.Login(ctx, &pb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		return
	}

	if !resp.Success {
		c.JSON(http.StatusUnauthorized, gin.H{"error": resp.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":       true,
		"message":       resp.Message,
		"token":         resp.Token,
		"refresh_token": resp.RefreshToken,
		"user": gin.H{
			"id":         resp.User.Id,
			"first_name": resp.User.FirstName,
			"last_name":  resp.User.LastName,
			"email":      resp.User.Email,
			"phone":      resp.User.Phone,
			"address":    resp.User.Address,
		},
	})
}

// GetUser retrieves user information
func (h *UserHandler) GetUser(c *gin.Context) {
	userID := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.userClient.GetUser(ctx, &pb.GetUserRequest{
		UserId: userID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	if !resp.Success {
		c.JSON(http.StatusNotFound, gin.H{"error": resp.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"user": gin.H{
			"id":         resp.User.Id,
			"first_name": resp.User.FirstName,
			"last_name":  resp.User.LastName,
			"email":      resp.User.Email,
			"phone":      resp.User.Phone,
			"address":    resp.User.Address,
			"created_at": resp.User.CreatedAt,
			"updated_at": resp.User.UpdatedAt,
		},
	})
}

// UpdateUser updates user information
func (h *UserHandler) UpdateUser(c *gin.Context) {
	userID := c.Param("id")

	var req struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Phone     string `json:"phone"`
		Address   string `json:"address"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.userClient.UpdateUser(ctx, &pb.UpdateUserRequest{
		UserId:    userID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Address:   req.Address,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	if !resp.Success {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": resp.Message,
		"user": gin.H{
			"id":         resp.User.Id,
			"first_name": resp.User.FirstName,
			"last_name":  resp.User.LastName,
			"email":      resp.User.Email,
			"phone":      resp.User.Phone,
			"address":    resp.User.Address,
		},
	})
}

// DeleteUser deletes a user
func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.userClient.DeleteUser(ctx, &pb.DeleteUserRequest{
		UserId: userID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	if !resp.Success {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": resp.Message,
	})
}

// VerifyToken verifies a JWT token
func (h *UserHandler) VerifyToken(c *gin.Context) {
	var req struct {
		Token string `json:"token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.userClient.VerifyToken(ctx, &pb.VerifyTokenRequest{
		Token: req.Token,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"valid":   resp.Valid,
		"user_id": resp.UserId,
		"message": resp.Message,
	})
}

// AuthMiddleware validates JWT token from Authorization header
func (h *UserHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Extract token from "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		token := parts[1]

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		resp, err := h.userClient.VerifyToken(ctx, &pb.VerifyTokenRequest{
			Token: token,
		})

		if err != nil || !resp.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Store user ID in context for handlers to use
		c.Set("user_id", resp.UserId)
		c.Next()
	}
}
