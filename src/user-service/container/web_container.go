package container

import (
	"assesement-test-MicroServices/grpc/pb"
	"assesement-test-MicroServices/src/user-service/config"
	"assesement-test-MicroServices/src/user-service/repository"
	"assesement-test-MicroServices/src/user-service/use_case"
	"fmt"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type WebContainer struct {
	Env        *config.EnvConfig
	UserDB     *config.DatabaseConfig
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
	UserDBConfig := config.NewDBConfig(envConfig)

	UserRepository := repository.NewUserRepository()
	repositoryContainer := NewRepositoryContainer(UserRepository)

	UserUseCase := use_case.NewUserUseCase(UserDBConfig, UserRepository)
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, UserUseCase)

	useCaseContainer := NewUseCaseContainer(UserUseCase)
	webContainer := &WebContainer{
		Env:        envConfig,
		UserDB:     UserDBConfig,
		Repository: repositoryContainer,
		UseCase:    useCaseContainer,
		Grpc:       grpcServer,
	}

	return webContainer
}
