package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tamahavik/tasker-api/pkg/dto"
	"github.com/tamahavik/tasker-api/pkg/models"
	"github.com/tamahavik/tasker-api/pkg/services"
	"github.com/tamahavik/tasker-api/pkg/utils"
	"net/http"
)

type UserController interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type userControllerImpl struct {
	service services.UserService
}

func NewUserController(s services.UserService) UserController {
	return &userControllerImpl{
		service: s,
	}
}

func (c userControllerImpl) Create(ctx *gin.Context) {
	var d dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&d); err != nil {
		utils.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := models.User{
		Username: d.Username,
		Email:    d.Email,
		Password: d.Password,
		Name:     d.Name,
	}

	createdUser := c.service.Create(user)
	utils.ResponseSuccess(ctx, http.StatusOK, createdUser)
}

func (c userControllerImpl) FindAll(ctx *gin.Context) {
	users := c.service.FindAll()
	utils.ResponseSuccess(ctx, http.StatusOK, users)
}
