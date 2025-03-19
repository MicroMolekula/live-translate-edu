package main

import (
	_ "github.com/live-translate-edu/docs"
	"github.com/live-translate-edu/internal/di"
	"github.com/live-translate-edu/internal/server"
	"log"
	"net/http"
	_ "net/http/pprof"
)

// @title           LiveTranslateEdu API
// @version         1.0
// @description     API приложения LiveTranslateEdu.

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apiKey  ApiKeyAuth
// @in header
// @name Authorization
func main() {
	di.InitContainer()
	go func() {
		http.ListenAndServe(":8888", nil)
	}()

	err := di.Container.Invoke(func(server *server.Server) {
		server.Run()
	})

	if err != nil {
		log.Fatal(err)
	}
}
