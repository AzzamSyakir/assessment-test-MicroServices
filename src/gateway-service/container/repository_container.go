package container

import (
	"assesement-test-MicroServices/src/gateway-service/repository"
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
