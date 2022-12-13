package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tamahavik/tasker-api/pkg/models"
	"github.com/tamahavik/tasker-api/pkg/services"
	"net/http"
)

type UserController interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type userControllerImpl struct {
	service services.UserService
}

func NewUserController(s services.UserService) *userControllerImpl {
	return &userControllerImpl{
		service: s,
	}
}

func (c userControllerImpl) Create(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdUser := c.service.Create(user)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   createdUser,
	})
}

func (c userControllerImpl) FindAll(ctx *gin.Context) {
	users := c.service.FindAll()
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   users,
	})
}
