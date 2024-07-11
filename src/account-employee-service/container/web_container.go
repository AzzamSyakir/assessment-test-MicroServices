package container

import (
	"assessment-test-MicroService/grpc/pb"
	"assessment-test-MicroService/src/user-service/config"
	"assessment-test-MicroService/src/user-service/repository"
	"assessment-test-MicroService/src/user-service/use_case"
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
	userDBConfig := config.NewUserDBConfig(envConfig)

	userRepository := repository.NewUserRepository()
	repositoryContainer := NewRepositoryContainer(userRepository)

	userUseCase := use_case.NewUserUseCase(userDBConfig, userRepository)
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userUseCase)

	useCaseContainer := NewUseCaseContainer(userUseCase)
	webContainer := &WebContainer{
		Env:        envConfig,
		UserDB:     userDBConfig,
		Repository: repositoryContainer,
		UseCase:    useCaseContainer,
		Grpc:       grpcServer,
	}

	return webContainer
}
