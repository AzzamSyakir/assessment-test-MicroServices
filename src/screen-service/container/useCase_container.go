package container

import (
	"assesement-test-MicroServices/src/screen-service/use_case"
)

type UseCaseContainer struct {
	Screen *use_case.ScreenUseCase
}

func NewUseCaseContainer(
	Screen *use_case.ScreenUseCase,

) *UseCaseContainer {
	useCaseContainer := &UseCaseContainer{
		Screen: Screen,
	}
	return useCaseContainer
}
