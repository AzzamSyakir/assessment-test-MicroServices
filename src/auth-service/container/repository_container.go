package container

import (
	"assesement-test-MicroServices/src/auth-service/repository"
)

type RepositoryContainer struct {
	Auth *repository.AuthRepository
}

func NewRepositoryContainer(
	Auth *repository.AuthRepository,

) *RepositoryContainer {
	repositoryContainer := &RepositoryContainer{
		Auth: Auth,
	}
	return repositoryContainer
}
