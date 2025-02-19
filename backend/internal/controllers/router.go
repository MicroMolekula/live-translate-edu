package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/live-translate-edu/internal/controllers/middleware"
)

type Router struct {
	auth            *AuthController
	room            *RoomController
	speechTranslate *SpeechTranslatorController
	user            *UserController
}

func NewRouter(
	auth *AuthController,
	room *RoomController,
	speechTranslate *SpeechTranslatorController,
	user *UserController) *Router {
	return &Router{
		auth:            auth,
		room:            room,
		speechTranslate: speechTranslate,
		user:            user,
	}
}

func (r *Router) InitRoutes(engine *gin.Engine) {
	engine.Use(middleware.CORSMiddleware())

	apiGroup := engine.Group("/api")
	{
		authRequiredGroup := apiGroup.Group("")
		authRequiredGroup.Use(r.auth.AuthRequiredMiddleware())
		{
			authRequiredGroup.GET("/connect", r.speechTranslate.Connect)
			authRequiredGroup.GET("/disconnect", r.speechTranslate.Disconnect)
			authRequiredGroup.GET("/me", r.auth.Me)
			authRequiredGroup.POST("/user/create", r.user.Create)
			authRequiredGroup.GET("/users", r.auth.Users)
		}
		apiGroup.POST("/auth", r.auth.Auth)
		apiGroup.POST("/token", r.room.GetJoinToken)
	}
}
