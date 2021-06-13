package accounts

import (
	"context"
	"github.com/sptGabriel/banking/app/domain/entities/accounts"
)

func (uc useCase) CreateAccount(ctx context.Context, acc accounts.Account) error {
	hasAccount, err := uc.repository.DoesExistByCPF(ctx, acc.CPF)
	if err != nil {
		return err
	}
	if hasAccount {
		return accounts.ErrAccountAlreadyExists
	}

	if acc.Secret.IsHashed() {
		return uc.repository.Create(ctx, &acc)
	}

	if err := acc.Secret.Encrypt(uc.hasher); err != nil {
		return err
	}

	return uc.repository.Create(ctx, &acc)
}
