package main

import (
	"github.com/live-translate-edu/internal/di"
	"github.com/live-translate-edu/internal/server"
	"log"
	"net/http"
	_ "net/http/pprof"
)

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
