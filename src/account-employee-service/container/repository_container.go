package container

import (
	"assessment-test-MicroService/src/user-service/repository"
)

type RepositoryContainer struct {
	User *repository.UserRepository
}

func NewRepositoryContainer(
	user *repository.UserRepository,

) *RepositoryContainer {
	repositoryContainer := &RepositoryContainer{
		User: user,
	}
	return repositoryContainer
}
