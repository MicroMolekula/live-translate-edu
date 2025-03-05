package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/live-translate-edu/internal/services/speech_translate"
	"net/http"
	"time"
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
	flagError := false
	go func() {
		defer func() {
			if err := recover(); err != nil {
				rc.speechTranslator.Stop()
				flagError = true
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"error":   err.(error).Error(),
					"message": "Ошибка распознования",
				})
			}
		}()
		rc.speechTranslator.SpeechTranslate("myroom")
	}()
	time.Sleep(2 * time.Second)
	if !flagError {
		ctx.JSON(200, gin.H{
			"success": true,
		})
	}
}

func (rc *SpeechTranslatorController) Disconnect(ctx *gin.Context) {
	go rc.speechTranslator.Stop()
	ctx.JSON(200, gin.H{
		"success": true,
	})
}
