package account

import (
	"context"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/entities"

	"github.com/sptGabriel/banking/app/infrastructure/database/postgres"
)

func (r Repository) UpdateBalance(ctx context.Context, account *entities.Account) error {
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