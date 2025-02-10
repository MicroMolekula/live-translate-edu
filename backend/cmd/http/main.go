package main

import (
	"backend/internal/configs"
	"backend/internal/controllers"
	"backend/internal/database"
	"fmt"
	"log"
)

func main() {
	configs.LoadConfig()
	database.Init()

	r := controllers.InitRouter()
	err := r.Run(fmt.Sprintf(":%d", configs.Cfg.ServerPort))
	if err != nil {
		log.Fatalf("Ошибка при запуске http сервера: %s", err)
	}
}
