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
	Office   string
	Password string
	Database string
}

type EnvConfig struct {
	App      *AppEnv
	OfficeDB *MongoEnv
}

func NewEnvConfig() *EnvConfig {
	envConfig := &EnvConfig{
		App: &AppEnv{
			Host: os.Getenv("GATEWAY_APP_HOST"),
			Port: os.Getenv("ROLE_SERVICES_PORT"),
		},
		OfficeDB: &MongoEnv{
			Host:     os.Getenv("MONGO_HOST"),
			Port:     os.Getenv("MONGO_PORT"),
			Office:   os.Getenv("MONGO_USER"),
			Password: os.Getenv("MONGO_PASSWORD"),
			Database: "db",
		},
	}
	return envConfig
}
