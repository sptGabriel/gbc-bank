package accounts

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/entities/accounts"
	"github.com/sptGabriel/banking/app/domain/vos"
)

func (r Repository) GetByID(ctx context.Context, id vos.AccountId) (*accounts.Account, error) {
	operation := "AccountRepository.GetById"
	query := "SELECT id, name, cpf, secret, balance, created_at FROM accounts WHERE id = $1"

	var account accounts.Account

	err := r.pool.QueryRow(ctx, query, id).Scan(
		&account.Id,
		&account.Name,
		&account.CPF,
		&account.Secret,
		&account.Balance,
		&account.CreatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrAccountNotFound
		}
		return nil, app.Err(operation, err)
	}

	return &account, nil
}
