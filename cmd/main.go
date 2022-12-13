package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tamahavik/tasker-api/pkg/controllers"
	"github.com/tamahavik/tasker-api/pkg/database"
	"github.com/tamahavik/tasker-api/pkg/repository"
	"github.com/tamahavik/tasker-api/pkg/services"
	"net/http"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())

	db, err := database.Connection()
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/users", userController.Create)

	r.GET("/users", userController.FindAll)
	err1 := r.Run(":8000")
	if err != nil {
		panic(err1)
	}
}
