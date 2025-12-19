package handler

import (
	"strconv"
	"time"

	"user-crud/internal/models"
	"user-crud/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.service.GetUser(c.Context(), id)
	if err != nil {
		return fiber.ErrNotFound
	}

	return c.JSON(user)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return fiber.ErrBadRequest
	}

	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return fiber.ErrBadRequest
	}

	err = h.service.CreateUser(c.Context(), req.Name, dob)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.service.ListUsers(c.Context())
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(users)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.ErrBadRequest
	}

	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return fiber.ErrBadRequest
	}

	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.service.UpdateUser(c.Context(), id, req.Name, dob)
	if err != nil {
		return fiber.ErrNotFound
	}

	return c.JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.service.DeleteUser(c.Context(), id); err != nil {
		return fiber.ErrNotFound
	}

	return c.SendStatus(fiber.StatusNoContent)
}
