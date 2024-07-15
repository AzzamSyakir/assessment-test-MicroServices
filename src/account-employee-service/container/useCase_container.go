package container

import (
	"assesement-test-MicroServices/src/account-employee-service/use_case"
)

type UseCaseContainer struct {
	Account *use_case.AccountUseCase
}

func NewUseCaseContainer(
	Account *use_case.AccountUseCase,

) *UseCaseContainer {
	useCaseContainer := &UseCaseContainer{
		Account: Account,
	}
	return useCaseContainer
}
