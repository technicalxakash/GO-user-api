package handler

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/akash/go-user-api/internal/logger"
	"github.com/akash/go-user-api/internal/models"
	"github.com/akash/go-user-api/internal/repository"
	"github.com/akash/go-user-api/internal/service"
)

type AuthHandler struct {
	repo     *repository.UserRepo
	validate *validator.Validate
}

func NewAuthHandler(r *repository.UserRepo) *AuthHandler {
	return &AuthHandler{
		repo:     r,
		validate: validator.New(),
	}
}

// Signup handles user registration
func (h *AuthHandler) Signup(c *fiber.Ctx) error {
	req := new(models.SignupRequest)

	if err := c.BodyParser(req); err != nil {
		return sendErrorResponse(c, 400, "INVALID_REQUEST", "Invalid request body")
	}

	// Validate request
	if err := h.validate.Struct(req); err != nil {
		return sendErrorResponse(c, 400, "VALIDATION_ERROR", err.Error())
	}

	// Check if email already exists
	existingUser, _ := h.repo.GetUserByEmail(c.Context(), req.Email)
	if existingUser != nil {
		return sendErrorResponse(c, 400, "EMAIL_EXISTS", "Email already registered")
	}

	// Hash password
	hashedPassword, err := service.HashPassword(req.Password)
	if err != nil {
		logger.Log.Error("Failed to hash password", zap.Error(err))
		return sendErrorResponse(c, 500, "INTERNAL_ERROR", "Failed to process request")
	}

	// Parse DOB
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return sendErrorResponse(c, 400, "INVALID_DATE", "Invalid date format")
	}

	// Create user
	err = h.repo.CreateUserWithAuth(c.Context(), req.Name, req.Email, hashedPassword, "user", dob)
	if err != nil {
		logger.Log.Error("Failed to create user", zap.Error(err))
		return sendErrorResponse(c, 500, "INTERNAL_ERROR", "Failed to create user")
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "User registered successfully",
		"email":   req.Email,
	})
}

// Login handles user authentication
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	req := new(models.LoginRequest)

	if err := c.BodyParser(req); err != nil {
		return sendErrorResponse(c, 400, "INVALID_REQUEST", "Invalid request body")
	}

	// Validate request
	if err := h.validate.Struct(req); err != nil {
		return sendErrorResponse(c, 400, "VALIDATION_ERROR", err.Error())
	}

	// Get user by email
	user, err := h.repo.GetUserByEmail(c.Context(), req.Email)
	if err != nil || user == nil {
		return sendErrorResponse(c, 401, "INVALID_CREDENTIALS", "Invalid email or password")
	}

	// Verify password
	if !service.VerifyPassword(user.PasswordHash, req.Password) {
		return sendErrorResponse(c, 401, "INVALID_CREDENTIALS", "Invalid email or password")
	}

	// Generate JWT token
	token, err := service.GenerateToken(user.ID, user.Email, user.Role, 24)
	if err != nil {
		logger.Log.Error("Failed to generate token", zap.Error(err))
		return sendErrorResponse(c, 500, "INTERNAL_ERROR", "Failed to generate token")
	}

	// Set cookie (optional)
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   false, // Set to true in production with HTTPS
		SameSite: "Strict",
	})

	return c.Status(200).JSON(models.AuthResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		Token: token,
	})
}

// GetProfile returns current user profile
func (h *AuthHandler) GetProfile(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(int32)
	if !ok {
		return sendErrorResponse(c, 401, "UNAUTHORIZED", "User not authenticated")
	}

	user, err := h.repo.GetUserByID(c.Context(), userID)
	if err != nil {
		return sendErrorResponse(c, 404, "NOT_FOUND", "User not found")
	}

	return c.JSON(models.UserProfile{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		DOB:   user.Dob.String(),
	})
}
