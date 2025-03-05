package controllers

import (
	"github.com/gin-gonic/gin"
)

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
