package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/live-translate-edu/internal/dto"
	"github.com/live-translate-edu/internal/services"
	"net/http"
)

type UtilsController struct {
	languageService *services.LanguageService
	groupService    *services.GroupService
}

func NewUtilsController(language *services.LanguageService, group *services.GroupService) *UtilsController {
	return &UtilsController{
		languageService: language,
		groupService:    group,
	}
}

func (uc *UtilsController) GetDataForCreateLesson(ctx *gin.Context) {
	languages, err := uc.languageService.GetAll()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err, "internal server error")
		return
	}
	groups, err := uc.groupService.GetGroups()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err, "internal server error")
		return
	}
	data := dto.CreateDataForFormLesson(languages, groups)
	ctx.JSON(http.StatusOK, data)
}
