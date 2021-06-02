package repositories

import (
	"context"
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/vos"
)

type AccountRepository interface {
	Create(ctx context.Context, account *entities.Account) error
	UpdateAccountBalance(ctx context.Context, account *entities.Account) error
	DoesAccountExistByCPF(ctx context.Context, cpf vos.CPF) (bool, error)
	GetByCPF(ctx context.Context, cpf vos.CPF) (*entities.Account, error)
	GetByID(ctx context.Context, id vos.AccountId) (*entities.Account, error)
	GetAccountBalance(ctx context.Context, accountId vos.AccountId) (*entities.Account, error)
	GetAll(ctx context.Context) ([]entities.Account, error)
}
