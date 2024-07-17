package container

import (
	"assesement-test-MicroServices/grpc/pb"
	"assesement-test-MicroServices/src/screen-service/config"
	"assesement-test-MicroServices/src/screen-service/repository"
	"assesement-test-MicroServices/src/screen-service/use_case"
	"fmt"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type WebContainer struct {
	Env        *config.EnvConfig
	ScreenDB   *config.DatabaseConfig
	Repository *RepositoryContainer
	UseCase    *UseCaseContainer
	Grpc       *grpc.Server
}

func NewWebContainer() *WebContainer {
	errEnvLoad := godotenv.Load()
	if errEnvLoad != nil {
		panic(fmt.Errorf("error loading .env file: %w", errEnvLoad))
	}

	envConfig := config.NewEnvConfig()
	ScreenDBConfig := config.NewDBConfig(envConfig)

	ScreenRepository := repository.NewScreenRepository()
	repositoryContainer := NewRepositoryContainer(ScreenRepository)

	ScreenUseCase := use_case.NewScreenUseCase(ScreenDBConfig, ScreenRepository)
	grpcServer := grpc.NewServer()
	pb.RegisterScreenServiceServer(grpcServer, ScreenUseCase)

	useCaseContainer := NewUseCaseContainer(ScreenUseCase)
	webContainer := &WebContainer{
		Env:        envConfig,
		ScreenDB:   ScreenDBConfig,
		Repository: repositoryContainer,
		UseCase:    useCaseContainer,
		Grpc:       grpcServer,
	}

	return webContainer
}
