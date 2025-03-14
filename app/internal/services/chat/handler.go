package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/live-translate-edu/internal/configs"
	"github.com/live-translate-edu/internal/dto"
	"github.com/live-translate-edu/internal/rabbitmq"
	"github.com/live-translate-edu/internal/services/speech_translate"
	"log"
)

type MessageHandlers struct {
	handlers []Handler
	cfg      *configs.Config
}

func NewMessageHandlers(cfg *configs.Config) *MessageHandlers {
	return &MessageHandlers{
		handlers: make([]Handler, 0),
		cfg:      cfg,
	}
}

func (mh *MessageHandlers) Add(h Handler) {
	mh.handlers = append(mh.handlers, h)
}

func (mh *MessageHandlers) HandleMessage(message *dto.MessageDto) {
	for _, h := range mh.handlers {
		h.messageHandle(message)
	}
}

type Handler interface {
	messageHandle(*dto.MessageDto)
}

type RabbitMQHandler struct {
	cfg *configs.Config
}

func NewRabbitMQHandler(cfg *configs.Config) *RabbitMQHandler {
	return &RabbitMQHandler{cfg: cfg}
}

// TODO Переписать Говно
func (rmqh *RabbitMQHandler) messageHandle(message *dto.MessageDto) {
	client, err := rabbitmq.NewRabbitMQ(rmqh.cfg)
	if err != nil {
		log.Println(err)
		return
	}
	defer client.Close()
	dataMessage, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return
	}
	if err = client.Publish("chat_messages", dataMessage); err != nil {
		log.Println(err)
		return
	}
}

type TranslateHandler struct {
	translateService *speech_translate.TranslateServ
}

func NewTranslateHandler(cfg *configs.Config) *TranslateHandler {
	translateService, err := speech_translate.NewTranslateServ(cfg)
	if err != nil {
		log.Println(fmt.Sprintf("Ошибка создания TranslateServ: %s", err.Error()))
	}
	return &TranslateHandler{
		translateService: translateService,
	}
}

func (th *TranslateHandler) messageHandle(message *dto.MessageDto) {
	translateText, err := th.translateService.TranslateText(context.Background(), message.Content, &dto.TranslateLanguagesDto{
		Source: "en",
		Target: "ru",
	})
	if err != nil {
		log.Println(fmt.Sprintf("Ошибка перевода сообщения {%s}: %v", message.Content, err))
		return
	}
	message.TranslateContent = translateText
}
