package accounts

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/entities/accounts"
	"github.com/sptGabriel/banking/app/domain/vos"
)

func (r Repository) GetByCPF(ctx context.Context, cpf vos.CPF) (*accounts.Account, error) {
	operation := "AccountRepository.GeyByCPF"
	query := "SELECT id, name, cpf, secret, balance, created_at FROM accounts WHERE cpf = $1"

	var account accounts.Account

	if err := r.pool.QueryRow(ctx, query, cpf).Scan(
		&account.Id,
		&account.Name,
		&account.CPF,
		&account.Secret,
		&account.Balance,
		&account.CreatedAt,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrAccountNotFound
		}
		return nil, app.Err(operation, err)
	}

	return &account, nil
}
