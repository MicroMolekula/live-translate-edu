package main

import (
	"backend/internal/configs"
	"backend/internal/controllers"
	"backend/internal/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	configs.LoadConfig()
	database.Init()

	recognizerController := controllers.NewRecognizerController()
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.POST("/api/token", controllers.GetJoinToken)
	r.GET("/api/connect", recognizerController.Connect)
	r.GET("/api/disconnect", recognizerController.Disconnect)
	r.GET("/api/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, world",
		})
	})
	err := r.Run(fmt.Sprintf(":%d", configs.Cfg.ServerPort))
	if err != nil {
		log.Fatalf("Ошибка при запуске http сервера: %s", err)
	}
}
