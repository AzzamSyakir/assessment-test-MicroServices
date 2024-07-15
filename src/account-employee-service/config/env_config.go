package config

import (
	"os"
)

type AppEnv struct {
	Host string
	Port string
}

type PostgresEnv struct {
	Host     string
	Port     string
	Account  string
	Password string
	Database string
}

type EnvConfig struct {
	App       *AppEnv
	AccountDB *PostgresEnv
}

func NewEnvConfig() *EnvConfig {
	envConfig := &EnvConfig{
		App: &AppEnv{
			Host: os.Getenv("GATEWAY_APP_HOST"),
			Port: os.Getenv("Account_SERVICES_PORT"),
		},
		AccountDB: &PostgresEnv{
			Host:     os.Getenv("MONGO_HOST"),
			Port:     os.Getenv("MONGO_PORT"),
			Account:  os.Getenv("MONGO_USER"),
			Password: os.Getenv("MONGO_PASSWORD"),
			Database: "db",
		},
	}
	return envConfig
}
