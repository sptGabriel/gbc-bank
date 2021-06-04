package repositories

import (
	"context"
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/vos"
)

type AccountRepository interface {
	Create(ctx context.Context, account *entities.Account) error
	UpdateBalance(ctx context.Context, account *entities.Account) error
	DoesExistByCPF(ctx context.Context, cpf vos.CPF) (bool, error)
	GetByID(ctx context.Context, cpf vos.AccountId) (*entities.Account, error)
	GetByCPF(ctx context.Context, cpf vos.CPF) (*entities.Account, error)
	GetBalance(ctx context.Context, accountId vos.AccountId) (*entities.Account, error)
	GetAll(ctx context.Context) ([]entities.Account, error)
}
