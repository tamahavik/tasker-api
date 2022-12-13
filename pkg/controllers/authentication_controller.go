package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tamahavik/tasker-api/pkg/services"
	"net/http"
)

type AuthenticationController interface {
	Login(ctx *gin.Context)
}

type authenticationControllerImpl struct {
	service services.AuthenticationService
}

func NewAuthentication(s services.AuthenticationService) *authenticationControllerImpl {
	return &authenticationControllerImpl{service: s}
}

func (a *authenticationControllerImpl) Login(ctx *gin.Context) {
	auth := ctx.Query("auth")
	password := ctx.Query("password")
	err := a.service.Login(auth, password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  401,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "success",
	})
}
