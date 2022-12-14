package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tamahavik/tasker-api/pkg/services"
	"github.com/tamahavik/tasker-api/pkg/utils"
	"net/http"
)

type AuthenticationController interface {
	Login(ctx *gin.Context)
}

type authenticationControllerImpl struct {
	service services.AuthenticationService
}

func NewAuthenticationController(s services.AuthenticationService) AuthenticationController {
	return &authenticationControllerImpl{service: s}
}

func (a *authenticationControllerImpl) Login(ctx *gin.Context) {
	auth := ctx.Query("auth")
	password := ctx.Query("password")
	err := a.service.Login(auth, password)
	if err != nil {
		utils.ResponseError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "success")
}
