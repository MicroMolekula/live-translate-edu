package configs

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Yandex  `yaml:"yandex"`
	Server  `yaml:"server"`
	LiveKit `yaml:"livekit"`
}

type Yandex struct {
	SpeechKitToken   string `yaml:"speech_kit_token"`
	TranslateToken   string `yaml:"translate_token"`
	AddressSpeechKit string `yaml:"address_speech_kit"`
	AddressTranslate string `yaml:"address_translate"`
}

type Server struct {
	ServerHost string `yaml:"server_host"`
	ServerPort int    `yaml:"server_port"`
}

type LiveKit struct {
	LiveKitApiUrl    string `yaml:"livekit_url"`
	LiveKitApiKey    string `yaml:"livekit_api_key"`
	LiveKitApiSecret string `yaml:"livekit_api_secret"`
	LiveKitIdentity  string `yaml:"livekit_identity"`
}

func NewConfig() *Config {
	return &Config{}
}

var Cfg *Config

func LoadConfig() {
	Cfg = NewConfig()
	yamlPath := os.Getenv("CONFIG_PATH")
	if yamlPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}
	yamlFile, err := os.ReadFile(yamlPath)
	if err != nil {
		log.Fatalf("Error load yaml file: %s", err)
	}
	err = yaml.Unmarshal(yamlFile, Cfg)
	if err != nil {
		log.Fatalf("Error parse yaml file: %s", err)
	}
}
