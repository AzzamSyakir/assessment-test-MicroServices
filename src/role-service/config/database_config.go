package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseConfig struct {
	RoleDB *mongoDB
}

type mongoDB struct {
	Connection *mongo.Client
}

func NewDBConfig(envConfig *EnvConfig) *DatabaseConfig {
	databaseConfig := &DatabaseConfig{
		RoleDB: NewDB(envConfig),
	}
	return databaseConfig
}

func NewDB(envConfig *EnvConfig) *mongoDB {
	var uri string
	if envConfig.RoleDB.Password == "" {
		uri = fmt.Sprintf(
			"mongodb://%s@%s:%s/%s",
			envConfig.RoleDB.Role,
			envConfig.RoleDB.Host,
			envConfig.RoleDB.Port,
			envConfig.RoleDB.Database,
		)
	} else {
		uri = fmt.Sprintf(
			"mongodb://%s:%s@%s:%s/%s",
			envConfig.RoleDB.Role,
			envConfig.RoleDB.Password,
			envConfig.RoleDB.Host,
			envConfig.RoleDB.Port,
			envConfig.RoleDB.Database,
		)
	}

	connection, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	RoleDB := &mongoDB{
		Connection: connection,
	}
	return RoleDB
}
