package config

import (
	"os"
)

type AppEnv struct {
	Host string
	Port string
}

type MongoEnv struct {
	Host     string
	Port     string
	Screen   string
	Password string
	Database string
}

type EnvConfig struct {
	App      *AppEnv
	ScreenDB *MongoEnv
}

func NewEnvConfig() *EnvConfig {
	envConfig := &EnvConfig{
		App: &AppEnv{
			Host: os.Getenv("GATEWAY_APP_HOST"),
			Port: os.Getenv("SCREEN_SERVICES_PORT"),
		},
		ScreenDB: &MongoEnv{
			Host:     os.Getenv("MONGO_HOST"),
			Port:     os.Getenv("MONGO_PORT"),
			Screen:   os.Getenv("MONGO_USER"),
			Password: os.Getenv("MONGO_PASSWORD"),
			Database: "db",
		},
	}
	return envConfig
}
