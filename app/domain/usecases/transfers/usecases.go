package transfers

import (
	"context"
	"github.com/sptGabriel/banking/app/domain/entities/accounts"
	"github.com/sptGabriel/banking/app/domain/entities/transfers"
	"github.com/sptGabriel/banking/app/domain/vos"
	"github.com/sptGabriel/banking/app/gateway/db/postgres"
)

type useCase struct {
	accountRepo   accounts.Repository
	transferRepo  transfers.Repository
	transactional postgres.Transactional
}

type UseCase interface {
	CreateTransfer(ctx context.Context, transfer transfers.Transfer) error
	GetTransfers(ctx context.Context, id vos.AccountId) ([]transfers.Transfer, error)
}

func NewUseCase(accRepo accounts.Repository, tr transfers.Repository, transactional postgres.Transactional) *useCase {
	return &useCase{accRepo, tr, transactional}
}
