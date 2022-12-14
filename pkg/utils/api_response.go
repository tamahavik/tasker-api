package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Status  string `json:"status,omitempty"`
	Data    any    `json:"data,omitempty"`
	Message any    `json:"message,omitempty"`
}

func ResponseSuccess(ctx *gin.Context, status int, data any) {
	var response = Response{
		Status: http.StatusText(status),
		Data:   data,
	}
	ctx.JSON(status, response)
}

func ResponseError(ctx *gin.Context, status int, message any) {
	var response = Response{
		Status:  http.StatusText(status),
		Message: message,
	}
	ctx.JSON(status, response)
}
