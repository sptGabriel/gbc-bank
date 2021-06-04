package account

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/entities"
)

func (r Repository) GetAll(ctx context.Context) ([]entities.Account, error) {
	operation := "AccountRepository.GetAll"
	query := `select id, name, cpf, balance, created_at from accounts`

	var accounts []entities.Account

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		if err == pgx.ErrNoRows {
			return accounts, nil
		}
		return nil, app.Err(operation, err)
	}

	for rows.Next() {
		var account entities.Account

		err := rows.Scan(
			&account.Id,
			&account.Name,
			&account.CPF,
			&account.Balance,
			&account.CreatedAt)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	if err := rows.Err(); err != nil {
		return nil, app.Err(operation, err)
	}

	return accounts, nil
}
