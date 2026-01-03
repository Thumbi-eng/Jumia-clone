package service

import (
	"errors"
	"time"

	"jumia-clone-backend/services/user-service/internal/models"
	"jumia-clone-backend/services/user-service/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your-secret-key-change-this-in-production") // TODO: Move to environment variable

// UserService defines business logic methods for user operations
type UserService interface {
	Register(firstName, lastName, email, password, phone string) (*models.User, error)
	Login(email, password string) (*models.User, string, string, error)
	GetUserByID(id string) (*models.User, error)
	UpdateUser(id, firstName, lastName, phone, address string) (*models.User, error)
	DeleteUser(id string) error
	VerifyToken(tokenString string) (string, error)
}

type userService struct {
	repo repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// Register creates a new user account
func (s *userService) Register(firstName, lastName, email, password, phone string) (*models.User, error) {
	// Check if user already exists
	existingUser, _ := s.repo.GetByEmail(email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  string(hashedPassword),
		Phone:     phone,
		IsActive:  true,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Login authenticates a user and returns JWT tokens
func (s *userService) Login(email, password string) (*models.User, string, string, error) {
	// Get user by email
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, "", "", errors.New("invalid email or password")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", "", errors.New("invalid email or password")
	}

	// Generate access token (expires in 24 hours)
	accessToken, err := generateToken(user.ID, 24*time.Hour)
	if err != nil {
		return nil, "", "", err
	}

	// Generate refresh token (expires in 7 days)
	refreshToken, err := generateToken(user.ID, 7*24*time.Hour)
	if err != nil {
		return nil, "", "", err
	}

	return user, accessToken, refreshToken, nil
}

// GetUserByID retrieves a user by ID
func (s *userService) GetUserByID(id string) (*models.User, error) {
	return s.repo.GetByID(id)
}

// UpdateUser updates user information
func (s *userService) UpdateUser(id, firstName, lastName, phone, address string) (*models.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields
	if firstName != "" {
		user.FirstName = firstName
	}
	if lastName != "" {
		user.LastName = lastName
	}
	if phone != "" {
		user.Phone = phone
	}
	if address != "" {
		user.Address = address
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser soft deletes a user
func (s *userService) DeleteUser(id string) error {
	return s.repo.Delete(id)
}

// VerifyToken verifies JWT token and returns user ID
func (s *userService) VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["user_id"].(string)
		if !ok {
			return "", errors.New("invalid token claims")
		}
		return userID, nil
	}

	return "", errors.New("invalid token")
}

// generateToken creates a JWT token for a user
func generateToken(userID string, expiration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(expiration).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
