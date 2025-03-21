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

// AddGroup Добавление новой группы студентов
//
// @Tags admin
// @Summary Добавление новой группы студентов
// @Description Добавление новой группы студентов
// @Security ApiKeyAuth
// @Param data body dto.Group true "данные о группе"
// @Success 201 {object} SuccessResponse
// @Failure 500 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 400 {object} ErrorResponse
// @Router /admin/groups/create [post]
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

// GetGroups
//
// @Summary Получение списка всех групп
// @Tags group
// @Security ApiKeyAuth
// @Success 200 {array} dto.Group
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /groups [get]
func (gac *GroupController) GetGroups(ctx *gin.Context) {
	groups, err := gac.groupService.GetGroups()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err, "failed to get groups")
		return
	}
	ctx.JSON(http.StatusOK, groups)
}

// AddUsersInGroup
//
// @Summary Добавление пользователя в группу
// @Tags admin
// @Security ApiKeyAuth
// @Param id path int true "Id группы"
// @Param q body controllers.AddUsersInGroup.userIds true "Массив id студентов"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /admin/groups/:id/users/add [post]
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

// GetUsers
//
// @Summary Получение всех студентов из группы
// @Tags admin
// @Security ApiKeyAuth
// @Param id path int true "Id группы"
// @Success {array} dto.UserDTO
// @Router /admin/groups/:id/users [get]
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
