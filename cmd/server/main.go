package main

import (
	"context"
	"fmt"
	"time"

	"user-crud/config"
	"user-crud/db/sqlc"
	"user-crud/internal/repository"
	"user-crud/internal/service"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	queries := sqlc.New(db)

	// userRepo := repository.NewUserRepository(queries)

	// user, err := userRepo.CreateUser(
	// 	context.Background(),
	// 	"Alice",
	// 	time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC),
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Inserted user via repository:", user)

	userRepo := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepo)

	err := userService.CreateUser(
		context.Background(),
		"Alice",
		time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("User created via service")

}
