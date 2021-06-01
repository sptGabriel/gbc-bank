package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/repositories"
	"github.com/sptGabriel/banking/app/domain/vos"
)

type accountRepository struct {
	conn *pgxpool.Pool
}

func (a accountRepository) DoesAccountExistByCPF(ctx context.Context, cpf vos.CPF) (bool, error) {
	panic("implement me")
}

func (a accountRepository) Create(ctx context.Context, account *entities.Account) error {
	panic("implement me")
}

func (a accountRepository) UpdateAccountBalance(ctx context.Context, accountId *vos.AccountId) error {
	panic("implement me")
}

func (a accountRepository) GetByCPF(ctx context.Context, cpf vos.CPF) (*entities.Account, error) {
	panic("implement me")
}

func (a accountRepository) GetAccountBalance(ctx context.Context, accountId vos.AccountId) (*entities.Account, error) {
	panic("implement me")
}

func (a accountRepository) GetAll(ctx context.Context) ([]entities.Account, error) {
	panic("implement me")
}

func NewAccountRepository(c *pgxpool.Pool) repositories.AccountRepository {
	return &accountRepository{c}
}
