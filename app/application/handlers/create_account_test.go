package handlers

import (
	"context"
	"errors"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/common/adapters"
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/repositories"
	"github.com/sptGabriel/banking/app/domain/vos"
	"github.com/sptGabriel/banking/app/tests/fakedata"
	"github.com/stretchr/testify/assert"
	"testing"
)



func TestCreateAccount(t *testing.T) {
	t.Run("Should create an account", func(t *testing.T) {
		repository := &repositories.AccountRepositoryMock{}
		hasher := adapters.NewBCryptAdapter(10)

		cmd := fakedata.FakeCreateAccountCMD()

		repository.CreateFunc = func(ctx context.Context, en *entities.Account) error {
			return nil
		}
		repository.DoesExistByCPFFunc = func(ctx context.Context, cpf vos.CPF) (bool, error) {
			return false, nil
		}

		handler := &createAccountHandler{
			hasher: hasher,
			repository: repository,
		}

		res, err := handler.Execute(context.Background(), cmd)

		assert.Nil(t, err)
		assert.Nil(t, res)
	})
	t.Run("Should return an error if handler.execute is called with invalid command", func(t *testing.T) {
		repository := &repositories.AccountRepositoryMock{}
		hasher := adapters.NewBCryptAdapter(10)

		badCmd := struct{}{}
		expectedError := app.Err("Handlers.CreateAccount", errors.New("invalid transfer command"))

		repository.CreateFunc = func(ctx context.Context, en *entities.Account) error {
			return nil
		}
		repository.DoesExistByCPFFunc = func(ctx context.Context, cpf vos.CPF) (bool, error) {
			return false, nil
		}

		handler := &createAccountHandler{
			hasher: hasher,
			repository: repository,
		}

		res, err := handler.Execute(context.Background(), badCmd)
		assert.Nil(t, res)
		assert.Equal(t, expectedError, err)
	})
	t.Run("Should return an already exists error", func(t *testing.T) {
		repository := &repositories.AccountRepositoryMock{}
		hasher := adapters.NewBCryptAdapter(10)

		cmd := fakedata.FakeCreateAccountCMD()

		repository.CreateFunc = func(ctx context.Context, en *entities.Account) error {
			return nil
		}
		repository.DoesExistByCPFFunc = func(ctx context.Context, cpf vos.CPF) (bool, error) {
			return true, nil
		}

		handler := &createAccountHandler{
			hasher: hasher,
			repository: repository,
		}

		res, err := handler.Execute(context.Background(), cmd)

		assert.Nil(t, res)
		assert.Equal(t, app.ErrAccountAlreadyExists, err)
	})
}
