package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tamahavik/tasker-api/pkg/services"
	"net/http"
)

type AuthenticationController interface {
	Login(c *gin.Context)
}

type authenticationControllerImpl struct {
	service services.AuthenticationService
}

func NewAuthentication(service services.AuthenticationService) *authenticationControllerImpl {
	return &authenticationControllerImpl{service: service}
}

func (a *authenticationControllerImpl) Login(c *gin.Context) {
	auth := c.Query("auth")
	password := c.Query("password")
	err := a.service.Login(auth, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  401,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "success",
	})
}
