package main

import (
	"context"
	"fmt"
	"time"

	"user-crud/config"
	"user-crud/db/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	queries := sqlc.New(db)

	dob := pgtype.Date{
		Time:  time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC),
		Valid: true,
	}

	// Insert user
	user, err := queries.CreateUser(context.Background(), sqlc.CreateUserParams{
		Name: "Alice",
		Dob:  dob,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Inserted user:", user)
}
