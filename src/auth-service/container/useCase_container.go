package container

import (
	"assesement-test-MicroServices/src/auth-service/use_case"
)

type UseCaseContainer struct {
	Auth *use_case.AuthUseCase
}

func NewUseCaseContainer(
	Auth *use_case.AuthUseCase,

) *UseCaseContainer {
	useCaseContainer := &UseCaseContainer{
		Auth: Auth,
	}
	return useCaseContainer
}
