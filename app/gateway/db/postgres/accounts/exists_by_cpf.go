package accounts

import (
	"context"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/vos"
)

func (r Repository) DoesExistByCPF(ctx context.Context, cpf vos.CPF) (bool, error) {
	operation := "AccountRepository.DoestExistsByCPF"
	query := `SELECT EXISTS(SELECT id FROM accounts WHERE cpf = $1)`
	accountExists := false
	if err := r.pool.QueryRow(ctx, query, cpf.String()).Scan(&accountExists); err != nil {
		return false, app.Err(operation, err)
	}
	return accountExists, nil
}
