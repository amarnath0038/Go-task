package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"user-crud/config"
	"user-crud/db/sqlc"
	"user-crud/internal/handler"
	"user-crud/internal/repository"
	"user-crud/internal/routes"
	"user-crud/internal/service"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	queries := sqlc.New(db)

	userRepo := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	app := fiber.New()

	routes.Register(app, userHandler)

	log.Fatal(app.Listen(":8080"))
}
