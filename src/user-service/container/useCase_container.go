package container

import (
	"assesement-test-MicroServices/src/user-service/use_case"
)

type UseCaseContainer struct {
	User *use_case.UserUseCase
}

func NewUseCaseContainer(
	User *use_case.UserUseCase,

) *UseCaseContainer {
	useCaseContainer := &UseCaseContainer{
		User: User,
	}
	return useCaseContainer
}
