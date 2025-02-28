package configs

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Yandex   `yaml:"yandex"`
	Server   `yaml:"server"`
	LiveKit  `yaml:"livekit"`
	Database `yaml:"database"`
	JWT      `yaml:"jwt"`
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

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Timezone string `yaml:"timezone"`
}

type JWT struct {
	Secret string `yaml:"secret"`
	TTL    int    `yaml:"ttl"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	yamlPath := os.Getenv("CONFIG_PATH")
	if yamlPath == "" {
		return nil, errors.New("config file environment variable not set")
	}
	yamlFile, err := os.ReadFile(yamlPath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error read config file: %s", err))
	}
	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error parse config file: %s", err))
	}
	return cfg, nil
}
