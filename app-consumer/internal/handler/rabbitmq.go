package handler

import (
	"app-consumer/internal/configs"
	"app-consumer/internal/rabbitmq"
	"fmt"
	"log"
)

type RabbitMQHandler struct {
	cfg *configs.Config
}

func NewRabbitMQHandler(cfg *configs.Config) *RabbitMQHandler {
	return &RabbitMQHandler{
		cfg: cfg,
	}
}

func (rh *RabbitMQHandler) Run() {
	rabbit, err := rabbitmq.NewRabbitMQ(rh.cfg)
	if err != nil {
		log.Fatal(err)
	}
	ch, err := rabbit.Consume("chat_messages")
	if err != nil {
		log.Fatal(err)
	}
	for msg := range ch {
		fmt.Printf("Сообщение из чата: %s\n", msg.Body)
	}
}
