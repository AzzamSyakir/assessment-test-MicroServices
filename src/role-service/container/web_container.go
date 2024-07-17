package container

import (
	"assesement-test-MicroServices/grpc/pb"
	"assesement-test-MicroServices/src/role-service/config"
	"assesement-test-MicroServices/src/role-service/repository"
	"assesement-test-MicroServices/src/role-service/use_case"
	"fmt"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type WebContainer struct {
	Env        *config.EnvConfig
	RoleDB     *config.DatabaseConfig
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
	RoleDBConfig := config.NewDBConfig(envConfig)

	RoleRepository := repository.NewRoleRepository()
	repositoryContainer := NewRepositoryContainer(RoleRepository)

	RoleUseCase := use_case.NewRoleUseCase(RoleDBConfig, RoleRepository)
	grpcServer := grpc.NewServer()
	pb.RegisterRoleServiceServer(grpcServer, RoleUseCase)

	useCaseContainer := NewUseCaseContainer(RoleUseCase)
	webContainer := &WebContainer{
		Env:        envConfig,
		RoleDB:     RoleDBConfig,
		Repository: repositoryContainer,
		UseCase:    useCaseContainer,
		Grpc:       grpcServer,
	}

	return webContainer
}
