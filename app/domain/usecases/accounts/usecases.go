package accounts

import (
	"context"
	"github.com/sptGabriel/banking/app/domain/entities/accounts"
	"github.com/sptGabriel/banking/app/domain/vos"
	"github.com/sptGabriel/banking/app/gateway/ports"
)

type useCase struct {
	repository accounts.Repository
	hasher     ports.Hash
}

type UseCase interface {
	CreateAccount(ctx context.Context, account accounts.Account) error
	GetBalance(ctx context.Context, id vos.AccountId) (accounts.Account, error)
	GetAll(ctx context.Context) ([]accounts.Account, error)
}

func NewUseCase(ar accounts.Repository, h ports.Hash) *useCase {
	return &useCase{ar, h}
}
