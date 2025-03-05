package configs

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	RabbitMQ
}

type RabbitMQ struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
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
