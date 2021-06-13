package accounts

import (
	"context"
	"github.com/sptGabriel/banking/app/domain/entities/accounts"
)

func (uc useCase) GetAll(ctx context.Context) ([]accounts.Account, error) {
	accounts, err := uc.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
