package handlers

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/commands"
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/repositories"
	"github.com/sptGabriel/banking/app/domain/vos"
	"github.com/sptGabriel/banking/app/infrastructure/database/postgres"
	"github.com/sptGabriel/banking/app/tests/fakedata"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMakeTransfer(t *testing.T) {
	transferRepository := &repositories.TransferRepositoryMock{}
	accountRepository := &repositories.AccountRepositoryMock{}
	transactional := &postgres.TransactionalMock{}

	fakeTransfer := fakedata.FakeTransfer()

	handler := &makeTransferHandler{
		accountRepo:   accountRepository,
		transferRepo:  transferRepository,
		transactional: transactional,
	}
	//cmd := commands.MakeTransferCommand{}

	t.Run("Should make a transfer", func(t *testing.T) {
		AccountOriginId := uuid.New()
		DestinationId := uuid.New()

		cmd := commands.NewMakeTransferCommand(AccountOriginId, DestinationId, fakeTransfer.Amount)


		transactional.ExecFunc = func(ctx context.Context, f func(context.Context) (interface{}, error)) (interface{}, error) {
			return nil, nil
		}

		res, err := handler.Execute(context.Background(), cmd)

		assert.Nil(t, err)
		assert.Nil(t, res)
	})

	t.Run("Should returns an error", func(t *testing.T) {
		AccountOriginId := uuid.New()
		DestinationId := uuid.New()

		cmd := commands.NewMakeTransferCommand(AccountOriginId, DestinationId, fakeTransfer.Amount)
		expectedErr := errors.New("failed to exec transactional")

		transactional.ExecFunc = func(ctx context.Context, f func(context.Context) (interface{}, error)) (interface{}, error) {
			return nil, expectedErr
		}

		res, err := handler.Execute(context.Background(), cmd)

		assert.Nil(t, res)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("Should return an error if handler.execute is called with invalid command", func(t *testing.T) {
		badCmd := struct{}{}

		expectedError := app.Err("Handlers.MakeTransfer", errors.New("invalid make transfer command"))

		res, err := handler.Execute(context.Background(), badCmd)
		assert.Nil(t, res)
		assert.Equal(t, reflect.TypeOf(&app.DomainError{}), reflect.TypeOf(err))
		assert.Equal(t, expectedError, err)
	})

	//make transfer . debit func

	t.Run("Should pass debit balance from origin account func", func(t *testing.T) {
		fakeAccount := fakedata.FakeAccount()

		accountRepository.GetByIDFunc = func(ctx context.Context, cpf vos.AccountId) (*entities.Account, error) {
			return fakeAccount, nil
		}
		accountRepository.UpdateBalanceFunc = func(ctx context.Context, account *entities.Account) error {
			return nil
		}

		err := handler.debitOriginAccount(context.Background(), fakeAccount.Id, fakeTransfer.Amount)

		assert.Nil(t, err)
	})
	t.Run("Should returns not found error origin account", func(t *testing.T) {
		fakeAccount := fakedata.FakeAccount()

		accountRepository.GetByIDFunc = func(ctx context.Context, cpf vos.AccountId) (*entities.Account, error) {
			return nil, app.ErrAccountNotFound
		}

		err := handler.debitOriginAccount(context.Background(), fakeAccount.Id, fakeTransfer.Amount)

		assert.Equal(t, app.ErrAccountNotFound, err)
	})

	t.Run("Should returns ErrInsufficientBalance error ", func(t *testing.T) {
		fakeAccount := fakedata.FakeAccount()

		accountRepository.GetByIDFunc = func(ctx context.Context, cpf vos.AccountId) (*entities.Account, error) {
			return fakeAccount, nil
		}

		accountRepository.UpdateBalanceFunc = func(ctx context.Context, account *entities.Account) error {
			return nil
		}
		err := handler.debitOriginAccount(context.Background(), fakeAccount.Id, 9999999)

		assert.Equal(t, app.ErrInsufficientBalance, err)

	})
	t.Run("Should returns error on update balance ", func(t *testing.T) {
		fakeAccount := fakedata.FakeAccount()

		expectedError := app.ErrBalanceUpdate

		accountRepository.GetByIDFunc = func(ctx context.Context, cpf vos.AccountId) (*entities.Account, error) {
			return fakeAccount, nil
		}
		accountRepository.UpdateBalanceFunc = func(ctx context.Context, account *entities.Account) error {
			return expectedError
		}

		err := handler.debitOriginAccount(context.Background(), fakeAccount.Id, fakeTransfer.Amount)

		assert.Equal(t, expectedError, err)
	})

	//make transfer . credit func

	t.Run("Should pass credit balance from destination account func", func(t *testing.T) {
		fakeAccount := fakedata.FakeAccount()

		accountRepository.GetByIDFunc = func(ctx context.Context, cpf vos.AccountId) (*entities.Account, error) {
			return fakeAccount, nil
		}
		accountRepository.UpdateBalanceFunc = func(ctx context.Context, account *entities.Account) error {
			return nil
		}

		err := handler.creditTargetAccount(context.Background(), fakeAccount.Id, fakeTransfer.Amount)

		assert.Nil(t, err)
	})
	t.Run("Should returns not found error target account", func(t *testing.T) {
		fakeAccount := fakedata.FakeAccount()

		accountRepository.GetByIDFunc = func(ctx context.Context, cpf vos.AccountId) (*entities.Account, error) {
			return nil, app.ErrAccountNotFound
		}

		err := handler.creditTargetAccount(context.Background(), fakeAccount.Id, fakeTransfer.Amount)

		assert.Equal(t, app.ErrAccountNotFound, err)
	})

	t.Run("Should returns ErrInvalidAmount error ", func(t *testing.T) {
		fakeAccount := fakedata.FakeAccount()

		accountRepository.GetByIDFunc = func(ctx context.Context, cpf vos.AccountId) (*entities.Account, error) {
			return fakeAccount, nil
		}

		accountRepository.UpdateBalanceFunc = func(ctx context.Context, account *entities.Account) error {
			return nil
		}
		err := handler.creditTargetAccount(context.Background(), fakeAccount.Id, 0)

		assert.Equal(t, app.ErrInvalidAmount, err)

	})
	t.Run("Should returns error on update balance ", func(t *testing.T) {
		fakeAccount := fakedata.FakeAccount()

		expectedError := app.ErrBalanceUpdate

		accountRepository.GetByIDFunc = func(ctx context.Context, cpf vos.AccountId) (*entities.Account, error) {
			return fakeAccount, nil
		}
		accountRepository.UpdateBalanceFunc = func(ctx context.Context, account *entities.Account) error {
			return expectedError
		}

		err := handler.creditTargetAccount(context.Background(), fakeAccount.Id, fakeTransfer.Amount)

		assert.Equal(t, expectedError, err)
	})

}
