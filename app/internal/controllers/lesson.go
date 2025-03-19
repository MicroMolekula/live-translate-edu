package controllers

import (
	"errors"
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
	user := ctx.Value("user").(*dto.UserDTO)
	if lessonRequest.TeacherId == 0 {
		lessonRequest.TeacherId = int(user.Id)
	}
	if err := lc.lessonService.CreateLesson(&lessonRequest); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err, "internal server error")
		return
	}
	newSuccessResponse(ctx, http.StatusCreated, "Урок успено создан", nil)
}

func (lc *LessonController) GetAll(ctx *gin.Context) {
	lessons, err := lc.lessonService.GetAll()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err, "internal server error")
		return
	}
	if len(lessons) == 0 {
		newErrorResponse(ctx, http.StatusNotFound, errors.New("not found"), "lesson not found")
		return
	}
	ctx.JSON(http.StatusOK, lessons)
}

func (lc *LessonController) GetByUser(ctx *gin.Context) {
	user := ctx.Value("user").(*dto.UserDTO)
	lessons, err := lc.lessonService.GetByUser(user)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err, "internal server error")
		return
	}
	if len(lessons) == 0 {
		newErrorResponse(ctx, http.StatusNotFound, errors.New("not found"), "lesson not found")
		return
	}
	ctx.JSON(http.StatusOK, lessons)
}
