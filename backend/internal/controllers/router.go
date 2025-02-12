package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/live-translate-edu/internal/controllers/middleware"
)

func InitRouter() *gin.Engine {

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	recognizerController := newRecognizerController()
	authController := newAuthController()
	userController := newUserController()

	apiGroup := r.Group("/api")
	{
		authRequiredGroup := apiGroup.Group("")
		authRequiredGroup.Use(authController.authRequiredMiddleware())
		{
			authRequiredGroup.GET("/connect", recognizerController.connect)
			authRequiredGroup.GET("/disconnect", recognizerController.disconnect)
			authRequiredGroup.GET("/me", authController.me)
			authRequiredGroup.POST("/token", getJoinToken)
		}
		apiGroup.POST("/auth", authController.auth)
		apiGroup.POST("/user/create", userController.create)
		apiGroup.GET("/users", authController.users)
	}

	return r
}
