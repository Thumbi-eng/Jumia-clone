package handler

import (
	"context"
	"time"

	"jumia-clone-backend/services/user-service/internal/service"
	pb "jumia-clone-backend/services/user-service/proto"
)

// UserServiceHandler implements the gRPC UserServiceServer
type UserServiceHandler struct {
	pb.UnimplementedUserServiceServer
	userService service.UserService
}

// NewUserServiceHandler creates a new user service handler
func NewUserServiceHandler(userService service.UserService) *UserServiceHandler {
	return &UserServiceHandler{
		userService: userService,
	}
}

// Register handles user registration
func (h *UserServiceHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user, err := h.userService.Register(
		req.FirstName,
		req.LastName,
		req.Email,
		req.Password,
		req.Phone,
	)
	if err != nil {
		return &pb.RegisterResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.RegisterResponse{
		UserId:  user.ID,
		Success: true,
		Message: "User registered successfully",
	}, nil
}

// Login handles user authentication
func (h *UserServiceHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, accessToken, refreshToken, err := h.userService.Login(req.Email, req.Password)
	if err != nil {
		return &pb.LoginResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.LoginResponse{
		UserId:       user.ID,
		Token:        accessToken,
		RefreshToken: refreshToken,
		Success:      true,
		Message:      "Login successful",
		User: &pb.UserData{
			Id:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
			Address:   user.Address,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		},
	}, nil
}

// GetUser retrieves user information
func (h *UserServiceHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := h.userService.GetUserByID(req.UserId)
	if err != nil {
		return &pb.GetUserResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.GetUserResponse{
		Success: true,
		Message: "User retrieved successfully",
		User: &pb.UserData{
			Id:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
			Address:   user.Address,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		},
	}, nil
}

// UpdateUser handles user information updates
func (h *UserServiceHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user, err := h.userService.UpdateUser(
		req.UserId,
		req.FirstName,
		req.LastName,
		req.Phone,
		req.Address,
	)
	if err != nil {
		return &pb.UpdateUserResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.UpdateUserResponse{
		Success: true,
		Message: "User updated successfully",
		User: &pb.UserData{
			Id:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
			Address:   user.Address,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		},
	}, nil
}

// DeleteUser handles user deletion
func (h *UserServiceHandler) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := h.userService.DeleteUser(req.UserId)
	if err != nil {
		return &pb.DeleteUserResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.DeleteUserResponse{
		Success: true,
		Message: "User deleted successfully",
	}, nil
}

// VerifyToken verifies JWT token
func (h *UserServiceHandler) VerifyToken(ctx context.Context, req *pb.VerifyTokenRequest) (*pb.VerifyTokenResponse, error) {
	userID, err := h.userService.VerifyToken(req.Token)
	if err != nil {
		return &pb.VerifyTokenResponse{
			Valid:   false,
			Message: err.Error(),
		}, nil
	}

	return &pb.VerifyTokenResponse{
		Valid:   true,
		UserId:  userID,
		Message: "Token is valid",
	}, nil
}
