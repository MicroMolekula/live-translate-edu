package controllers

import (
	"backend/internal/configs"
	"backend/internal/services/recognize"
	"github.com/gin-gonic/gin"
)

type SpeechTranslatorController struct {
	speechTranslator *recognize.SpeechTranslator
}

func NewRecognizerController() *SpeechTranslatorController {
	return &SpeechTranslatorController{
		recognize.NewSpeechTranslator(
			configs.Cfg.LiveKitApiUrl,
			configs.Cfg.LiveKitApiKey,
			configs.Cfg.LiveKitApiSecret,
			configs.Cfg.LiveKitApiSecret,
		),
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
