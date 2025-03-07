package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/live-translate-edu/internal/configs"
	"github.com/live-translate-edu/internal/dto"
	"github.com/live-translate-edu/internal/services/chat"
	"log"
	"net/http"
)

type ChatController struct {
	upgrader websocket.Upgrader
	hub      *chat.Hub
	cfg      *configs.Config
}

func NewChatController(cfg *configs.Config) *ChatController {
	handlers := chat.NewMessageHandlers(cfg)
	handlers.Add(chat.NewTranslateHandler(cfg))
	handlers.Add(chat.NewRabbitMQHandler(cfg))
	hub := chat.NewHub(handlers)
	go hub.Run()
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return &ChatController{
		upgrader: upgrader,
		hub:      hub,
	}
}

func (c *ChatController) Connect(ctx *gin.Context) {
	c.upgrader.Subprotocols = []string{"auth", ctx.Value("jwt").(string)}
	user := ctx.Value("user").(*dto.UserDTO)
	room := ctx.Param("room")
	wsConn, err := c.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		newErrorResponse(ctx, http.StatusInternalServerError, err, "Ошибка подключения websocket")
	}
	client := chat.NewClient(room, wsConn, c.hub, user)

	c.hub.AddRegisterClient(client)
	go client.Write()
	go client.Read()
}

func (c *ChatController) GetAllUsers(ctx *gin.Context) {
	room := ctx.Param("room")
	users, err := c.hub.GetUsers(room)
	if err != nil {
		newErrorResponse(ctx, http.StatusNotFound, err, "Нет такой комнаты в данный момент")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}
