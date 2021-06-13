package accounts

import (
	"context"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/entities/accounts"
)

func (r Repository) GetAll(ctx context.Context) ([]accounts.Account, error) {
	const operation = "Repository.Accounts.GetAll"

	var query = `select id, name, cpf, balance, created_at from accounts`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, app.Err(operation, err)
	}

	defer rows.Close()

	var accountList []accounts.Account
	for rows.Next() {
		var account accounts.Account
		err := rows.Scan(
			&account.Id,
			&account.Name,
			&account.CPF,
			&account.Balance,
			&account.CreatedAt,
		)
		if err != nil {
			return nil, app.Err(operation, err)
		}
		accountList = append(accountList, account)
	}

	if err := rows.Err(); err != nil {
		return nil, app.Err(operation, err)
	}

	return accountList, nil
}
