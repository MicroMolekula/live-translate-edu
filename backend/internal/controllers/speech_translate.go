package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/live-translate-edu/internal/services/speech_translate"
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
	go rc.speechTranslator.SpeechTranslate("myroom")
	ctx.JSON(200, gin.H{
		"success": true,
	})
}

func (rc *SpeechTranslatorController) Disconnect(ctx *gin.Context) {
	go rc.speechTranslator.Stop()
	ctx.JSON(200, gin.H{
		"success": true,
	})
}
