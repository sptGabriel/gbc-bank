package transfers

import (
	"context"
	"errors"
	"github.com/sptGabriel/banking/app/domain/entities/transfers"
	"github.com/sptGabriel/banking/app/domain/vos"
)

var (
	ErrAccountOriginNotFound = errors.New("origin account not found")
	ErrAccountDestinationNotFound = errors.New("destination account not found")
)

func (uc useCase) CreateTransfer(ctx context.Context, transfer transfers.Transfer) error {
	_, err := uc.transactional.Exec(ctx, func(txCtx context.Context) (interface{}, error) {
		if err := uc.debitOriginAccount(txCtx, transfer.AccountOriginId, transfer.Amount); err != nil {
			return nil, err
		}
		if err := uc.creditTargetAccount(txCtx, transfer.AccountDestinationId, transfer.Amount); err != nil {
			return nil, err
		}
		err := uc.transferRepo.Create(txCtx, &transfer)
		return nil, err
	})
	return err
}

func (uc *useCase) debitOriginAccount(ctx context.Context, id vos.AccountId, amount int) error {
	account, err := uc.accountRepo.GetByID(ctx, id)
	if err != nil {
		return ErrAccountOriginNotFound
	}

	if err := account.DebitAmount(amount); err != nil {
		return err
	}

	return uc.accountRepo.UpdateBalance(ctx, account)
}

func (uc *useCase) creditTargetAccount(ctx context.Context, id vos.AccountId, amount int) error {
	account, err := uc.accountRepo.GetByID(ctx, id)

	if err != nil {
		return ErrAccountDestinationNotFound
	}

	if err := account.CreditAmount(amount); err != nil {
		return err
	}

	return uc.accountRepo.UpdateBalance(ctx, account)
}
