package main

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"user-crud/config"
	"user-crud/db/sqlc"
	"user-crud/internal/handler"
	"user-crud/internal/logger"
	"user-crud/internal/middleware"
	"user-crud/internal/repository"
	"user-crud/internal/routes"
	"user-crud/internal/service"
)

func main() {

	logger.Init()
	defer logger.Log.Sync()
	logger.Log.Info("starting user-crud server")

	db := config.ConnectDB()
	defer db.Close()

	queries := sqlc.New(db)

	userRepo := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	app := fiber.New()

	app.Use(middleware.RequestID())

	routes.Register(app, userHandler)

	//log.Fatal(app.Listen(":8080"))

	if err := app.Listen(":8080"); err != nil {
		logger.Log.Fatal("failed to start server", zap.Error(err))
	}

}
