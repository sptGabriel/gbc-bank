package accounts

import (
	"context"
	"github.com/sptGabriel/banking/app/domain/vos"
)

type Repository interface {
	Create(ctx context.Context, account *Account) error
	UpdateBalance(ctx context.Context, account *Account) error
	DoesExistByCPF(ctx context.Context, cpf vos.CPF) (bool, error)
	GetByID(ctx context.Context, cpf vos.AccountId) (*Account, error)
	GetByCPF(ctx context.Context, cpf vos.CPF) (*Account, error)
	GetAll(ctx context.Context) ([]Account, error)
}
