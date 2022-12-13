package controllers

import "github.com/gin-gonic/gin"

type UserController interface {
	Create(c *gin.Context)
	FindAll(c *gin.Context)
}
