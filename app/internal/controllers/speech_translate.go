package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/live-translate-edu/internal/services/speech_translate"
	"net/http"
)

type SpeechTranslatorController struct {
	speechTranslator *speech_translate.SpeechTranslator
}

func NewRecognizerController(speechTranslator *speech_translate.SpeechTranslator) *SpeechTranslatorController {
	return &SpeechTranslatorController{
		speechTranslator: speechTranslator,
	}
}

func (rc *SpeechTranslatorController) Connect(ctx *gin.Context) {
	type Request struct {
		RoomTitle string `form:"room"`
	}
	var room Request
	if err := ctx.BindQuery(&room); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	go rc.speechTranslator.SpeechTranslate(room.RoomTitle)
	ctx.JSON(200, gin.H{
		"success": true,
	})
}

func (rc *SpeechTranslatorController) Disconnect(ctx *gin.Context) {
	type Request struct {
		RoomTitle string `form:"room"`
	}
	var room Request
	if err := ctx.BindQuery(&room); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	go rc.speechTranslator.Stop(room.RoomTitle)
	ctx.JSON(200, gin.H{
		"success": true,
	})
}
