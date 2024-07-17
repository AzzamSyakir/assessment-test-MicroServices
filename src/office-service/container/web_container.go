package container

import (
	"assesement-test-MicroServices/grpc/pb"
	"assesement-test-MicroServices/src/office-service/config"
	"assesement-test-MicroServices/src/office-service/repository"
	"assesement-test-MicroServices/src/office-service/use_case"
	"fmt"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type WebContainer struct {
	Env        *config.EnvConfig
	OfficeDB   *config.DatabaseConfig
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
	OfficeDBConfig := config.NewDBConfig(envConfig)

	OfficeRepository := repository.NewOfficeRepository()
	repositoryContainer := NewRepositoryContainer(OfficeRepository)

	OfficeUseCase := use_case.NewOfficeUseCase(OfficeDBConfig, OfficeRepository)
	grpcServer := grpc.NewServer()
	pb.RegisterOfficeServiceServer(grpcServer, OfficeUseCase)

	useCaseContainer := NewUseCaseContainer(OfficeUseCase)
	webContainer := &WebContainer{
		Env:        envConfig,
		OfficeDB:   OfficeDBConfig,
		Repository: repositoryContainer,
		UseCase:    useCaseContainer,
		Grpc:       grpcServer,
	}

	return webContainer
}
