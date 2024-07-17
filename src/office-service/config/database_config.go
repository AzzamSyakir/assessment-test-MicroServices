package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseConfig struct {
	OfficeDB *mongoDB
}

type mongoDB struct {
	Connection *mongo.Client
}

func NewDBConfig(envConfig *EnvConfig) *DatabaseConfig {
	databaseConfig := &DatabaseConfig{
		OfficeDB: NewDB(envConfig),
	}
	return databaseConfig
}

func NewDB(envConfig *EnvConfig) *mongoDB {
	var uri string
	if envConfig.OfficeDB.Password == "" {
		uri = fmt.Sprintf(
			"mongodb://%s@%s:%s/%s",
			envConfig.OfficeDB.Office,
			envConfig.OfficeDB.Host,
			envConfig.OfficeDB.Port,
			envConfig.OfficeDB.Database,
		)
	} else {
		uri = fmt.Sprintf(
			"mongodb://%s:%s@%s:%s/%s",
			envConfig.OfficeDB.Office,
			envConfig.OfficeDB.Password,
			envConfig.OfficeDB.Host,
			envConfig.OfficeDB.Port,
			envConfig.OfficeDB.Database,
		)
	}

	connection, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	OfficeDB := &mongoDB{
		Connection: connection,
	}
	return OfficeDB
}
