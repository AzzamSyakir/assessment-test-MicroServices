package config

import (
	"os"
)

type AppEnv struct {
	Host        string
	AuthPort    string
	AccountPort string
	RolePort    string
	OfficePort  string
	ScreenPort  string
}

type MongoEnv struct {
	Host     string
	Port     string
	Auth     string
	Password string
	Database string
}

type EnvConfig struct {
	App    *AppEnv
	AuthDB *MongoEnv
}

func NewEnvConfig() *EnvConfig {
	envConfig := &EnvConfig{
		App: &AppEnv{
			Host:        os.Getenv("GATEWAY_APP_HOST"),
			AuthPort:    os.Getenv("AUTH_SERVICES_PORT"),
			AccountPort: os.Getenv("ACCOUNT_SERVICES_PORT"),
			RolePort:    os.Getenv("ROLE_SERVICES_PORT"),
			OfficePort:  os.Getenv("OFFICE_SERVICES_PORT"),
			ScreenPort:  os.Getenv("SCREEN_SERVICES_PORT"),
		},
		AuthDB: &MongoEnv{
			Host:     os.Getenv("MONGO_HOST"),
			Port:     os.Getenv("MONGO_PORT"),
			Auth:     os.Getenv("MONGO_USER"),
			Password: os.Getenv("MONGO_PASSWORD"),
			Database: "db",
		},
	}
	return envConfig
}
