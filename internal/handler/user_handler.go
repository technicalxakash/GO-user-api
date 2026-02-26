package handler

import (
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/akash/go-user-api/internal/logger"
	"github.com/akash/go-user-api/internal/models"
	"github.com/akash/go-user-api/internal/repository"
	"github.com/akash/go-user-api/internal/service"
)

type UserHandler struct {
	repo     *repository.UserRepo
	validate *validator.Validate
}

func NewUserHandler(r *repository.UserRepo) *UserHandler {
	return &UserHandler{
		repo:     r,
		validate: validator.New(),
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {

	req := new(models.CreateUserRequest)

	if err := c.BodyParser(req); err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.validate.Struct(req); err != nil {
		return fiber.ErrBadRequest
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return fiber.ErrBadRequest
	}

	err = h.repo.Create(c.Context(), req.Name, dob)
	if err != nil {
		logger.Log.Error("Failed to create user", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	return c.Status(201).JSON(req)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	user, err := h.repo.Get(c.Context(), int32(id))
	if err != nil {
		return fiber.ErrNotFound
	}

	age := service.CalculateAge(user.Dob)

	return c.JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob,
		"age":  age,
	})
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {

	users, err := h.repo.List(c.Context())
	if err != nil {
		logger.Log.Error("Failed to list users", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	result := make([]fiber.Map, 0)

	for _, u := range users {
		result = append(result, fiber.Map{
			"id":   u.ID,
			"name": u.Name,
			"dob":  u.Dob,
			"age":  service.CalculateAge(u.Dob),
		})
	}

	return c.JSON(result)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	req := new(models.CreateUserRequest)
	c.BodyParser(req)

	dob, _ := time.Parse("2006-01-02", req.DOB)

	err := h.repo.Update(c.Context(), int32(id), req.Name, dob)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(req)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	err := h.repo.Delete(c.Context(), int32(id))
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(204)
}
