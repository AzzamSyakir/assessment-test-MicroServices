package container

import (
	"assesement-test-MicroServices/src/office-service/use_case"
)

type UseCaseContainer struct {
	Office *use_case.OfficeUseCase
}

func NewUseCaseContainer(
	Office *use_case.OfficeUseCase,

) *UseCaseContainer {
	useCaseContainer := &UseCaseContainer{
		Office: Office,
	}
	return useCaseContainer
}
