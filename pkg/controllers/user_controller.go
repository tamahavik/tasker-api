package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tamahavik/tasker-api/pkg/models"
	"github.com/tamahavik/tasker-api/pkg/services"
	"net/http"
)

type UserController interface {
	Create(c *gin.Context)
	FindAll(c *gin.Context)
}

type userControllerImpl struct {
	services.UserService
}

func NewUserController(userService services.UserService) *userControllerImpl {
	return &userControllerImpl{
		UserService: userService,
	}
}

func (controller userControllerImpl) Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdUser := controller.UserService.Create(user)
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   createdUser,
	})
}

func (controller userControllerImpl) FindAll(c *gin.Context) {
	users := controller.UserService.FindAll()
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   users,
	})
}
