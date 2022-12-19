package config

import (
	"github.com/gin-gonic/gin"
	"github.com/tamahavik/tasker-api/pkg/injection"
	"gorm.io/gorm"
	"net/http"
)

type Route interface {
	InitRoute()
}

type route struct {
	e  *gin.Engine
	db *gorm.DB
}

func NewRoutes(e *gin.Engine, db *gorm.DB) Route {
	return &route{
		e:  e,
		db: db,
	}
}

func (r route) InitRoute() {
	userController := injection.InstanceUser(r.db)
	authController := injection.InstanceAuthentication(r.db)

	r.e.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.e.POST("/auth/login", authController.Login)
	r.e.POST("/auth/register", authController.Register)

	r.e.POST("/users", userController.Create)
	r.e.GET("/users", userController.FindAll)
}
