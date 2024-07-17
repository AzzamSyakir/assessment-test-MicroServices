package container

import (
	"assesement-test-MicroServices/src/screen-service/repository"
)

type RepositoryContainer struct {
	Screen *repository.ScreenRepository
}

func NewRepositoryContainer(
	Screen *repository.ScreenRepository,

) *RepositoryContainer {
	repositoryContainer := &RepositoryContainer{
		Screen: Screen,
	}
	return repositoryContainer
}
