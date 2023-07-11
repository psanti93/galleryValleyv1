package main

import (
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/psanti93/galleryValleyv1/models"
)

func main() {

	cfg := models.DefaultPostgresConfig()

	db, err := models.Open(cfg)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}

	userService := models.UserService{
		DB: db,
	}

	user, err := userService.Create("paul@santiago3.com", "test1")

	if err != nil {
		fmt.Errorf("Creating user: %v", err)
	}

	fmt.Println("User in DB: ", user)

}
