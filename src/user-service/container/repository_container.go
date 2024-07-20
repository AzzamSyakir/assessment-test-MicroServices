package container

import (
	"assesement-test-MicroServices/src/user-service/repository"
)

type RepositoryContainer struct {
	User *repository.UserRepository
}

func NewRepositoryContainer(
	User *repository.UserRepository,

) *RepositoryContainer {
	repositoryContainer := &RepositoryContainer{
		User: User,
	}
	return repositoryContainer
}
