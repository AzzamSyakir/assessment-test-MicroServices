package container

import (
	"assesement-test-MicroServices/src/role-service/use_case"
)

type UseCaseContainer struct {
	Role *use_case.RoleUseCase
}

func NewUseCaseContainer(
	Role *use_case.RoleUseCase,

) *UseCaseContainer {
	useCaseContainer := &UseCaseContainer{
		Role: Role,
	}
	return useCaseContainer
}
