package controllers

import (
	"backend/internal/services/room"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserIdentity struct {
	Name string `json:"name"`
	Room string `json:"room"`
}

func GetJoinToken(ctx *gin.Context) {
	var userIdentity UserIdentity
	if err := ctx.Bind(&userIdentity); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":       true,
			"error_message": err.Error(),
		})
	}
	token, err := room.GenerateJoinToken(
		userIdentity.Room,
		userIdentity.Name,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":       true,
			"error_message": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"token":   token,
	})
}
