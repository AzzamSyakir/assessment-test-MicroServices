package container

import (
	"assesement-test-MicroServices/grpc/pb"
	"assesement-test-MicroServices/src/gateway-service/config"
	"assesement-test-MicroServices/src/gateway-service/repository"
	"assesement-test-MicroServices/src/gateway-service/use_case"
	"fmt"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type WebContainer struct {
	Env        *config.EnvConfig
	AccountDB  *config.DatabaseConfig
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
	AccountDBConfig := config.NewDBConfig(envConfig)

	AccountRepository := repository.NewAccountRepository()
	repositoryContainer := NewRepositoryContainer(AccountRepository)

	AccountUseCase := use_case.NewAccountUseCase(AccountDBConfig, AccountRepository)
	grpcServer := grpc.NewServer()
	pb.RegisterAccountServiceServer(grpcServer, AccountUseCase)

	useCaseContainer := NewUseCaseContainer(AccountUseCase)
	webContainer := &WebContainer{
		Env:        envConfig,
		AccountDB:  AccountDBConfig,
		Repository: repositoryContainer,
		UseCase:    useCaseContainer,
		Grpc:       grpcServer,
	}

	return webContainer
}
