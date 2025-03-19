package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/live-translate-edu/internal/dto"
	"github.com/live-translate-edu/internal/services"
	"net/http"
)

type LanguageController struct {
	languageService *services.LanguageService
}

func NewLanguageController(language *services.LanguageService) *LanguageController {
	return &LanguageController{languageService: language}
}

func (lc *LanguageController) Create(ctx *gin.Context) {
	type languageRequest struct {
		Title string `json:"title" binding:"required"`
		Code  string `json:"code" binding:"required"`
	}
	var request languageRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err, "bad request")
		return
	}
	if err := lc.languageService.Create(request.Title, request.Code); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err, "internal server error")
		return
	}
	newSuccessResponse(ctx, http.StatusCreated, "Язык успешно добавлен", nil)
}

func (lc *LanguageController) GetAll(ctx *gin.Context) {
	var languagesResponse []dto.Language
	languagesResponse, err := lc.languageService.GetAll()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err, "internal server error")
		return
	}
	ctx.JSON(http.StatusOK, languagesResponse)
}
