package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/live-translate-edu/internal/controllers/middleware"
)

func InitRouter() *gin.Engine {

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	initRecognizerController(r)
	r.POST("/api/token", getJoinToken)

	return r
}

func initRecognizerController(r *gin.Engine) {
	recognizerController := newRecognizerController()

	r.GET("/api/connect", recognizerController.connect)
	r.GET("/api/disconnect", recognizerController.disconnect)
}
