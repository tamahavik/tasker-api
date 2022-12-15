package main

import (
	"fmt"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/tamahavik/tasker-api/pkg/database"
	"github.com/tamahavik/tasker-api/pkg/injection"
	"net/http"
)

func main() {
	config := injection.InstanceConfig()
	err := config.ReadConfiguration()
	if err != nil {
		panic(err)
	}

	gin.SetMode(viper.GetString("MODE"))
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(helmet.Default())
	r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedExtensions([]string{".pdf", ".mp4"})))

	conn := database.NewConnection()
	db, err := conn.GetConnection()
	if err != nil {
		panic(err)
	}

	userController := injection.InstanceUser(db)
	authController := injection.InstanceAuthentication(db)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/auth/login", authController.Login)

	r.POST("/users", userController.Create)
	r.GET("/users", userController.FindAll)
	err1 := r.Run(fmt.Sprintf(":%s", viper.GetString("PORT")))
	if err != nil {
		panic(err1)
	}
}
