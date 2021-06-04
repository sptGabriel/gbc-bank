package account

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/vos"
)

func (r Repository) GetBalance(ctx context.Context, accId vos.AccountId) (*entities.Account, error) {
	operation := "AccountRepository.GetBalance"
	query := "SELECT balance FROM accounts WHERE id = $1"

	var account entities.Account

	if err := r.pool.QueryRow(ctx, query, accId).Scan(
		&account.Id,
		&account.Name,
		&account.CPF,
		&account.Secret,
		&account.Balance,
		&account.CreatedAt,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, app.ErrAccountNotFound
		}
		return nil, app.Err(operation, err)
	}

	return &account, nil
}