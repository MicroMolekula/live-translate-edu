package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/live-translate-edu/internal/dto"
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
		newErrorResponse(ctx, http.StatusBadRequest, err, "Bad request")
		return
	}
	token, err := rc.roomService.GenerateJoinToken(
		userIdentity.Room,
		userIdentity.Name,
	)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err, "Internal server error")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"token":   token,
	})
}

func (rc *RoomController) GetRoomTokenForUser(ctx *gin.Context) {
	type room struct {
		Title string `form:"room"`
	}
	var request room
	if err := ctx.BindQuery(&request); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err, "Некорректное тело запроса")
		return
	}
	user := ctx.Value("user").(*dto.UserDTO)
	token, err := rc.roomService.GetRoomTokenForUser(user, request.Title)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err, "Internal server error")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"token":   token,
	})
}
