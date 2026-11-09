package config

import (
	"os"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env				string		`yaml:"env" env:"ENV" env-default:"local"`
	HTTPServer		HTTPServer	`yaml:"http_server"`
	Log				LogConfig	`yaml:"log"`
	Database   		Database   	`yaml:"database"`
}

type HTTPServer struct {
	Address		string		`yaml:"address" env:"HTTP_ADDRESS" env-default:"localhost:8080"`
}

type LogConfig struct {
	Level string `yaml:"level" env:"LOG_LEVEL" env-default:"debug"` // debug, info, warn, error
}

type Database struct {
	DSN string `yaml:"dsn" env:"DATABASE_DSN" env-default:"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"`
}

func Must() *Config {
	var cfg Config

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config/config.yaml"
	}

	if _, err := os.Stat(configPath); err == nil {
		if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
			log.Fatalf("failed to read config file: %s", err)
		}
	} else {
		if err := cleanenv.ReadEnv(&cfg); err != nil {
			log.Fatalf("failed to read env variables: %s", err)
		}
	}

	return &cfg
}