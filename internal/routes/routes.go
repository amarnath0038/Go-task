package routes

import (
	"user-crud/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, h *handler.UserHandler) {
	app.Get("/users/:id", h.GetUser)
	app.Get("/users", h.ListUsers)
	app.Post("/users", h.CreateUser)
	app.Put("/users/:id", h.UpdateUser)
	app.Delete("/users/:id", h.DeleteUser)
}
