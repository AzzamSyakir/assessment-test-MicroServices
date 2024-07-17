package container

import (
	"assesement-test-MicroServices/src/role-service/repository"
)

type RepositoryContainer struct {
	Account *repository.AccountRepository
}

func NewRepositoryContainer(
	Account *repository.AccountRepository,

) *RepositoryContainer {
	repositoryContainer := &RepositoryContainer{
		Account: Account,
	}
	return repositoryContainer
}
