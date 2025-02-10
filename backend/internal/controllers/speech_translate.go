package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/live-translate-edu/internal/configs"
	"github.com/live-translate-edu/internal/services/recognize"
)

type SpeechTranslatorController struct {
	speechTranslator *recognize.SpeechTranslator
}

func newRecognizerController() *SpeechTranslatorController {
	return &SpeechTranslatorController{
		recognize.NewSpeechTranslator(
			configs.Cfg.LiveKitApiUrl,
			configs.Cfg.LiveKitApiKey,
			configs.Cfg.LiveKitApiSecret,
			configs.Cfg.LiveKitApiSecret,
		),
	}
}

func (rc *SpeechTranslatorController) connect(ctx *gin.Context) {
	go rc.speechTranslator.SpeechTranslate("myroom")
	ctx.JSON(200, gin.H{
		"success": true,
	})
}

func (rc *SpeechTranslatorController) disconnect(ctx *gin.Context) {
	go rc.speechTranslator.Stop()
	ctx.JSON(200, gin.H{
		"success": true,
	})
}
