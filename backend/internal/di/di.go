package di

import (
	"github.com/live-translate-edu/internal/configs"
	"github.com/live-translate-edu/internal/controllers"
	"github.com/live-translate-edu/internal/database"
	"github.com/live-translate-edu/internal/repository"
	"github.com/live-translate-edu/internal/server"
	"github.com/live-translate-edu/internal/services"
	"github.com/live-translate-edu/internal/services/speech_translate"
	"go.uber.org/dig"
	"log"
)

var Container *dig.Container

func InitContainer() {
	Container = dig.New()

	for _, general := range getGeneral() {
		if err := Container.Provide(general); err != nil {
			log.Fatal(err)
		}
	}

	for _, serv := range getServices() {
		if err := Container.Provide(serv); err != nil {
			log.Fatal(err)
		}
	}

	for _, rep := range getRepositories() {
		if err := Container.Provide(rep); err != nil {
			log.Fatal(err)
		}
	}

	for _, cont := range getControllers() {
		if err := Container.Provide(cont); err != nil {
			log.Fatal(err)
		}
	}
}

func getGeneral() []interface{} {
	return []interface{}{
		configs.NewConfig,
		database.NewDB,
		server.NewServer,
		controllers.NewRouter,
	}
}

func getServices() []interface{} {
	return []interface{}{
		services.NewAuthService,
		services.NewJwtService,
		services.NewRoomService,
		services.NewUserService,
		speech_translate.NewSpeechTranslator,
		speech_translate.NewTranslateServ,
		speech_translate.NewRecognizer,
	}
}

func getRepositories() []interface{} {
	return []interface{}{
		repository.NewUserRepository,
	}
}

func getControllers() []interface{} {
	return []interface{}{
		controllers.NewRoomController,
		controllers.NewAuthController,
		controllers.NewUserController,
		controllers.NewRecognizerController,
	}
}
