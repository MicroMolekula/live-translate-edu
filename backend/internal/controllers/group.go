package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/live-translate-edu/internal/dto"
	"github.com/live-translate-edu/internal/services"
	"net/http"
	"strconv"
)

type GroupController struct {
	groupService *services.GroupService
}

func NewGroupController(groupService *services.GroupService) *GroupController {
	return &GroupController{groupService: groupService}
}

func (gac *GroupController) AddGroup(ctx *gin.Context) {
	var groupRequest *dto.Group
	if err := ctx.ShouldBindJSON(&groupRequest); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, errors.New("bad request"), "invalid request body")
		return
	}
	if err := gac.groupService.AddNewGroup(groupRequest); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err, "failed to add new group")
		return
	}
	newSuccessResponse(ctx, http.StatusCreated, "Группа успешно добавлена", nil)
}

func (gac *GroupController) GetGroups(ctx *gin.Context) {
	groups, err := gac.groupService.GetGroups()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err, "failed to get groups")
		return
	}
	ctx.JSON(http.StatusOK, groups)
}

func (gac *GroupController) AddUsersInGroup(ctx *gin.Context) {
	type userIds struct {
		Ids []int `json:"users_ids"`
	}
	var request userIds
	if err := ctx.ShouldBindJSON(&request); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err, "invalid request body")
		return
	}
	fmt.Println(request)
	groupId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	if err := gac.groupService.AddUsersInGroup(groupId, request.Ids); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err, "failed to add users to group")
		return
	}
	newSuccessResponse(ctx, http.StatusOK, "Пользователи добавлены в группу", nil)
}

func (gac *GroupController) GetUsers(ctx *gin.Context) {
	groupId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	users, err := gac.groupService.GetUsersByGroupId(groupId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err, "failed to get users")
		return
	}
	if users == nil {
		users = []*dto.UserDTO{}
	}
	ctx.JSON(http.StatusOK, users)
}
