package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tamahavik/tasker-api/pkg/dto"
	"github.com/tamahavik/tasker-api/pkg/services"
	"github.com/tamahavik/tasker-api/pkg/utils"
	"net/http"
)

type AuthenticationController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authenticationControllerImpl struct {
	service services.AuthenticationService
}

func NewAuthenticationController(s services.AuthenticationService) AuthenticationController {
	return &authenticationControllerImpl{service: s}
}

func (a *authenticationControllerImpl) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	err := a.service.Login(&req)
	if err != nil {
		utils.ResponseError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "success")
}

func (a *authenticationControllerImpl) Register(ctx *gin.Context) {
	var req dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	newUser, err := a.service.Register(&req)
	if err != nil {
		utils.ResponseError(ctx, http.StatusBadRequest, err.Error())
	}
	utils.ResponseSuccess(ctx, http.StatusOK, newUser)
}
