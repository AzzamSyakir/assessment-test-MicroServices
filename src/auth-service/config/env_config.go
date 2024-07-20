package config

import (
	"os"
)

type AppEnv struct {
	Host        string
	AccountHost string
	OfficeHost  string
	RoleHost    string
	ScreenHost  string
	UserHost    string
	AuthPort    string
	AccountPort string
	RolePort    string
	OfficePort  string
	ScreenPort  string
	UserPort    string
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
			AccountHost: os.Getenv("ACCOUNT_HOST"),
			OfficeHost:  os.Getenv("OFFICE_HOST"),
			RoleHost:    os.Getenv("ROLE_HOST"),
			ScreenHost:  os.Getenv("SCREEN_HOST"),
			UserHost:    os.Getenv("USER_HOST"),
			AuthPort:    os.Getenv("AUTH_SERVICES_PORT"),
			AccountPort: os.Getenv("ACCOUNT_SERVICES_PORT"),
			RolePort:    os.Getenv("ROLE_SERVICES_PORT"),
			OfficePort:  os.Getenv("OFFICE_SERVICES_PORT"),
			ScreenPort:  os.Getenv("SCREEN_SERVICES_PORT"),
			UserPort:    os.Getenv("USER_PORT"),
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
