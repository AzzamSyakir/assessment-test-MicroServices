package container

import (
	"assesement-test-MicroServices/src/role-service/repository"
)

type RepositoryContainer struct {
	Role *repository.RoleRepository
}

func NewRepositoryContainer(
	Role *repository.RoleRepository,

) *RepositoryContainer {
	repositoryContainer := &RepositoryContainer{
		Role: Role,
	}
	return repositoryContainer
}
