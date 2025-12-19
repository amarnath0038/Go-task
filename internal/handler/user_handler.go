package handler

import (
	"strconv"
	"time"

	"user-crud/internal/logger"
	"user-crud/internal/models"
	"user-crud/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserHandler struct {
	service  *service.UserService
	validate *validator.Validate
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		service:  s,
		validate: validator.New(),
	}
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		logger.Log.Warn("invalid user id", zap.String("id", c.Params("id")))
		return fiber.ErrBadRequest
	}

	user, err := h.service.GetUser(c.Context(), id)
	if err != nil {
		logger.Log.Warn("user not found", zap.Int64("id", id))
		return fiber.ErrNotFound
	}
	return c.JSON(user)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		logger.Log.Warn("failed to parse request body", zap.Error(err))
		return fiber.ErrBadRequest
	}

	if err := h.validate.Struct(req); err != nil {
		logger.Log.Warn("validation failed", zap.Error(err))
		return fiber.ErrBadRequest
	}

	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		logger.Log.Warn("invalid dob format", zap.String("dob", req.Dob))
		return fiber.ErrBadRequest
	}

	if err := h.service.CreateUser(c.Context(), req.Name, dob); err != nil {
		logger.Log.Error("failed to create user", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	logger.Log.Info("user created",
		zap.String("name", req.Name),
		zap.String("dob", req.Dob),
	)

	return c.SendStatus(fiber.StatusCreated)
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.service.ListUsers(c.Context())
	if err != nil {
		logger.Log.Error("failed to list users", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	return c.JSON(users)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		logger.Log.Warn("invalid user id", zap.String("id", c.Params("id")))
		return fiber.ErrBadRequest
	}

	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		logger.Log.Warn("failed to parse request body", zap.Error(err))
		return fiber.ErrBadRequest
	}

	if err := h.validate.Struct(req); err != nil {
		logger.Log.Warn("validation failed", zap.Error(err))
		return fiber.ErrBadRequest
	}

	dob, err := time.Parse("2025-12-19", req.Dob)
	if err != nil {
		logger.Log.Warn("invalid dob format", zap.String("dob", req.Dob))
		return fiber.ErrBadRequest
	}

	user, err := h.service.UpdateUser(c.Context(), id, req.Name, dob)
	if err != nil {
		logger.Log.Warn("failed to update user", zap.Int64("id", id))
		return fiber.ErrNotFound
	}

	logger.Log.Info("user updated",
		zap.Int64("id", id),
		zap.String("name", req.Name),
	)

	return c.JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		logger.Log.Warn("invalid user id", zap.String("id", c.Params("id")))
		return fiber.ErrBadRequest
	}

	if err := h.service.DeleteUser(c.Context(), id); err != nil {
		logger.Log.Warn("failed to delete user", zap.Int64("id", id))
		return fiber.ErrNotFound
	}

	logger.Log.Info("user deleted", zap.Int64("id", id))
	return c.SendStatus(fiber.StatusNoContent)
}
