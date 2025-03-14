package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/live-translate-edu/internal/controllers/middleware"
	"github.com/live-translate-edu/internal/utils/roles"
)

type Router struct {
	auth            *AuthController
	room            *RoomController
	speechTranslate *SpeechTranslatorController
	user            *UserController
	chat            *ChatController
	group           *GroupController
	language        *LanguageController
	lesson          *LessonController
}

func NewRouter(
	auth *AuthController,
	room *RoomController,
	speechTranslate *SpeechTranslatorController,
	user *UserController,
	chat *ChatController,
	group *GroupController,
	language *LanguageController,
	lesson *LessonController) *Router {
	return &Router{
		auth:            auth,
		room:            room,
		speechTranslate: speechTranslate,
		user:            user,
		chat:            chat,
		group:           group,
		language:        language,
		lesson:          lesson,
	}
}

func (r *Router) InitRoutes(engine *gin.Engine) {
	engine.Use(middleware.CORSMiddleware())

	apiGroup := engine.Group("/api")
	{
		authRequiredGroup := apiGroup.Group("")
		authRequiredGroup.Use(r.auth.AuthRequiredMiddleware())
		{
			adminGroup := authRequiredGroup.Group("/admin")
			adminGroup.Use(middleware.RoleMiddleware([]string{roles.Admin}))
			{
				adminGroup.POST("/groups/create", r.group.AddGroup)
				adminGroup.GET("/groups", r.group.GetGroups)
				adminGroup.POST("/user/create", r.user.Create)
				adminGroup.GET("/users", r.auth.Users)
				adminGroup.GET("/groups/:id/users", r.group.GetUsers)
				adminGroup.POST("/groups/:id/users/add", r.group.AddUsersInGroup)
				adminGroup.POST("/language/create", r.language.Create)
			}
			authRequiredGroup.GET("/connect", r.speechTranslate.Connect)
			authRequiredGroup.GET("/disconnect", r.speechTranslate.Disconnect)
			authRequiredGroup.GET("/me", r.auth.Me)
			authRequiredGroup.GET("/user/room_token", r.room.GetRoomTokenForUser)
			authRequiredGroup.GET("/chat/connect/:room", r.chat.Connect)
			authRequiredGroup.GET("/chat/:room/users", r.chat.GetAllUsers)
			authRequiredGroup.POST("/lesson/create", r.lesson.CreateLesson)
		}
		apiGroup.POST("/auth", r.auth.Auth)
		apiGroup.POST("/token", r.room.GetJoinToken)
	}
}
