package controllers

import (
	"backend/internal/controllers/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	recognizerController := newRecognizerController()

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/token", getJoinToken)
	r.GET("/api/connect", recognizerController.connect)
	r.GET("/api/disconnect", recognizerController.disconnect)
	return r
}
