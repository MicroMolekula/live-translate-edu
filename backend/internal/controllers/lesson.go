package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/live-translate-edu/internal/dto"
	"github.com/live-translate-edu/internal/services"
	"net/http"
)

type LessonController struct {
	lessonService *services.LessonService
}

func NewLessonController(lessonService *services.LessonService) *LessonController {
	return &LessonController{
		lessonService: lessonService,
	}
}

func (lc *LessonController) CreateLesson(ctx *gin.Context) {
	var lessonRequest dto.LessonCreate
	if err := ctx.ShouldBindJSON(&lessonRequest); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err, "bad request")
		return
	}
	if err := lc.lessonService.CreateLesson(&lessonRequest); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err, "internal server error")
		return
	}
	newSuccessResponse(ctx, http.StatusCreated, "Урок успено создан", nil)
}
