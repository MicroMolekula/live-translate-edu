package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/live-translate-edu/internal/dto"
	"github.com/live-translate-edu/internal/services"
	"net/http"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Create(ctx *gin.Context) {
	var request *dto.UserCreateDTO
	if err := ctx.ShouldBind(&request); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err, "Bad request")
		return
	}
	err := uc.userService.CreateUser(request)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err, "Internal server error")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"user":    request,
	})
}
