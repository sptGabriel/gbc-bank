package accounts

import (
	"github.com/sptGabriel/banking/app/domain/entities/accounts"
	"github.com/sptGabriel/banking/app/gateway/ports"
)

var (
	mockedUseCase    UseCase
	mockedRepository *accounts.RepositoryMock
	mockedHasher     *ports.HashMock
)

func setupUseCaseTest() {
	mockedRepository = &accounts.RepositoryMock{}
	mockedHasher = &ports.HashMock{}
	mockedUseCase = useCase{
		mockedRepository,
		mockedHasher,
	}
}
