package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseConfig struct {
	ScreenDB *mongoDB
}

type mongoDB struct {
	Connection *mongo.Client
}

func NewDBConfig(envConfig *EnvConfig) *DatabaseConfig {
	databaseConfig := &DatabaseConfig{
		ScreenDB: NewDB(envConfig),
	}
	return databaseConfig
}

func NewDB(envConfig *EnvConfig) *mongoDB {
	var uri string
	if envConfig.ScreenDB.Password == "" {
		uri = fmt.Sprintf(
			"mongodb://%s@%s:%s/%s",
			envConfig.ScreenDB.Screen,
			envConfig.ScreenDB.Host,
			envConfig.ScreenDB.Port,
			envConfig.ScreenDB.Database,
		)
	} else {
		uri = fmt.Sprintf(
			"mongodb://%s:%s@%s:%s/%s",
			envConfig.ScreenDB.Screen,
			envConfig.ScreenDB.Password,
			envConfig.ScreenDB.Host,
			envConfig.ScreenDB.Port,
			envConfig.ScreenDB.Database,
		)
	}
	fmt.Println("uri", uri)
	connection, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	ScreenDB := &mongoDB{
		Connection: connection,
	}
	return ScreenDB
}
