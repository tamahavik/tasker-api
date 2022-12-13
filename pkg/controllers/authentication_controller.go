package controllers

import "github.com/gin-gonic/gin"

type AuthenticationController interface {
	Login(c *gin.Context)
}
