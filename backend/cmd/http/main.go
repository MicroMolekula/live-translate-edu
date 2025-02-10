package main

import (
	"fmt"
	"github.com/live-translate-edu/internal/configs"
	"github.com/live-translate-edu/internal/controllers"
	"github.com/live-translate-edu/internal/database"
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
