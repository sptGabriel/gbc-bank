package accounts

import (
	"context"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/entities/accounts"
	"github.com/sptGabriel/banking/app/gateway/db/postgres"
)

func (r Repository) UpdateBalance(ctx context.Context, account *accounts.Account) error {
	operation := "AccountRepository.UpdateBalance"
	query := "UPDATE accounts SET balance = $1 WHERE id = $2"
	if _, err := postgres.GetConnFromCtx(ctx, r.pool).Exec(ctx, query,
		account.Balance,
		account.Id,
	); err != nil {
		return app.Err(operation, err)
	}

	return nil
}
