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
