package transfers

import (
	"github.com/sptGabriel/banking/app/domain/entities/accounts"
	"github.com/sptGabriel/banking/app/domain/entities/transfers"
	"github.com/sptGabriel/banking/app/gateway/db/postgres"
)

var (
	mockedUseCase            UseCase
	mockedAccRepository      *accounts.RepositoryMock
	mockedTransferRepository *transfers.RepositoryMock
	mockedTransactional      *postgres.TransactionalMock
)

func setupUseCaseTest() {
	mockedTransferRepository = &transfers.RepositoryMock{}
	mockedAccRepository = &accounts.RepositoryMock{}
	mockedTransactional = &postgres.TransactionalMock{}
	mockedUseCase = useCase{
		mockedAccRepository,
		mockedTransferRepository,
		mockedTransactional,
	}
}
