package account

import (
	"context"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/entities"
)

func (r Repository) Create(ctx context.Context, account *entities.Account) error {
	operation := "AccountRepository.Create"
	query := `
		INSERT INTO
			accounts (id, name, cpf, secret, balance)
		VALUES ($1, $2, $3, $4, $5)
		`
	_ ,err := r.pool.Exec(ctx, query,
		account.Id,
		account.Name,
		account.CPF,
		account.Secret,
		account.Balance,
	)

	if err == nil {
		return nil
	}

	return app.Err(operation, err)
}