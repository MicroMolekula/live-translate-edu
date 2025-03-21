package controllers

import (
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type TokenResponse struct {
	Token   string `json:"token"`
	Success bool   `json:"success"`
}

func newErrorResponse(ctx *gin.Context, status int, err error, message string) {
	ctx.JSON(status, gin.H{
		"error":   err.Error(),
		"success": false,
		"message": message,
	})
}

func newSuccessResponse(ctx *gin.Context, status int, message string, data map[string]interface{}) {
	if data == nil {
		ctx.JSON(status, gin.H{
			"success": true,
			"message": message,
		})
		return
	}
	ctx.JSON(status, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}
