package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tamahavik/tasker-api/pkg/database"
	"github.com/tamahavik/tasker-api/pkg/models"
)

func main() {
	db, err := database.Connection()
	if err != nil {
		panic(err)
	}


	var user models.User = models.User{
		Id:        uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username:  "hatama",
		Email:     "priyatama.havik@gmail.com",
		Password:  "123456",
		Name:      "Havik Priyatama",
		Age:       30,
	}

	db.Create(&user)
	fmt.Printf("success inserted with id %s", user.Id)
}
