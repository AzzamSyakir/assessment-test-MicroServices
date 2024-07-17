package container

import (
	"assesement-test-MicroServices/src/office-service/repository"
)

type RepositoryContainer struct {
	Office *repository.OfficeRepository
}

func NewRepositoryContainer(
	Office *repository.OfficeRepository,

) *RepositoryContainer {
	repositoryContainer := &RepositoryContainer{
		Office: Office,
	}
	return repositoryContainer
}
