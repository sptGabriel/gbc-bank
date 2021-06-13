package accounts

import (
	"context"
	"github.com/sptGabriel/banking/app/domain/entities/accounts"
	"github.com/sptGabriel/banking/app/domain/vos"
)

func (uc useCase) GetBalance(ctx context.Context, id vos.AccountId) (*accounts.Account, error) {
	account, err := uc.repository.GetByID(ctx, id)
	return account, err
}
