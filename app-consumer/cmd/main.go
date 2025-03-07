package main

import (
	"app-consumer/internal/configs"
	"app-consumer/internal/handler"
	"fmt"
	"log"
)

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\nStart RabbitMQ Handler ///")
	handlerRabbit := handler.NewRabbitMQHandler(cfg)
	handlerRabbit.Run()
}
