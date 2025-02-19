package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/live-translate-edu/internal/services"
	"net/http"
)

type UserIdentity struct {
	Name string `json:"name"`
	Room string `json:"room"`
}

type RoomController struct {
	roomService *services.RoomService
}

func NewRoomController(roomService *services.RoomService) *RoomController {
	return &RoomController{
		roomService: roomService,
	}
}

func (rc *RoomController) GetJoinToken(ctx *gin.Context) {
	var userIdentity UserIdentity
	if err := ctx.Bind(&userIdentity); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":       true,
			"error_message": err.Error(),
		})
	}
	token, err := rc.roomService.GenerateJoinToken(
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
